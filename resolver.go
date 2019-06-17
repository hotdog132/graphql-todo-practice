package graphql_todo_practice

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {
	todo := &Todo{
		Text: input.Text,
		User: &User{
			ID:   input.UserID,
			Name: fmt.Sprintf("user-%d", input.UserID),
		},
	}

	if err := storeUser(todo.User); err != nil {
		return nil, err
	}

	if err := storeTodoEvent(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*Todo, error) {
	todos, err := getTodoEvents()
	if err != nil {
		return nil, err
	}

	m := make(map[int]*User)

	for _, t := range todos {
		if _, ok := m[t.UserID]; !ok {
			u, err := getUserByID(t.UserID)
			if err != nil {
				return nil, err
			}
			m[t.UserID] = u
		}

		t.User = m[t.UserID]
	}

	return todos, nil
}

func storeUser(u *User) error {
	url := "http://user-service:7070/users"

	m := make(map[string]interface{})
	m["id"] = u.ID
	m["name"] = u.Name

	body, err := json.Marshal(m)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("store user got error")
	}

	return nil
}

func storeTodoEvent(t *Todo) error {
	url := "http://event-service:6060/events"

	m := make(map[string]interface{})
	m["user_id"] = t.User.ID
	m["text"] = t.Text

	body, err := json.Marshal(m)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("store event got error")
	}

	return nil
}

func getUserByID(userID int) (*User, error) {
	url := fmt.Sprintf("http://user-service:7070/users/%d", userID)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	sitemap, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var u *User
	if err := json.Unmarshal(sitemap, &u); err != nil {
		return nil, err
	}

	return u, nil
}

func getTodoEvents() ([]*Todo, error) {
	res, err := http.Get("http://event-service:6060/events")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	sitemap, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var todos []*Todo
	if err := json.Unmarshal(sitemap, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}
