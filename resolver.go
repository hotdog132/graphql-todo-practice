package graphql_todo_practice

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	todos []*Todo
}

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
		ID:   rand.Int(),
		User: &User{
			ID:   input.UserID,
			Name: fmt.Sprintf("user-%s", input.UserID),
		},
	}
	r.todos = append(r.todos, todo)
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

func getUserByID(userID int) (*User, error) {
	url := fmt.Sprintf("http://localhost:7070/users/%d", userID)
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
	res, err := http.Get("http://localhost:6060/events")
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
