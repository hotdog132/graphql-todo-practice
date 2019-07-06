package delivery

import (
	"net/http"
	"strconv"

	"github.com/hotdog132/graphql-todo-practice/user-service/models"
	"github.com/hotdog132/graphql-todo-practice/user-service/user"
	"github.com/labstack/echo"
)

type HttpUserHandler struct {
	UserUsecase user.Usecase
}

func NewHttpUserHandler(e *echo.Echo, uu user.Usecase) {
	uh := &HttpUserHandler{uu}

	e.GET("/users/:id", uh.FetchUserByID)
	e.GET("/users", uh.FetchUsers)
	e.POST("/users", uh.StoreUser)
}

func (uh *HttpUserHandler) StoreUser(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	if err := uh.UserUsecase.Store(u); err != nil {
		r := make(map[string]string)
		r["status"] = "Existed"
		c.JSON(http.StatusOK, r)
	}

	return c.JSON(http.StatusOK, u)
}

func (uh *HttpUserHandler) FetchUsers(c echo.Context) error {
	users, err := uh.UserUsecase.FetchAll()
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func (uh *HttpUserHandler) FetchUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	u, err := uh.UserUsecase.Fetch(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, u)
}
