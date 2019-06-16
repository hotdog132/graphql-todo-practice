package delivery

import (
	"net/http"
	"strconv"

	"github.com/hotdog132/graphql-todo-practice/event-service/event"
	"github.com/hotdog132/graphql-todo-practice/event-service/models"
	"github.com/labstack/echo"
)

type HttpEventHandler struct {
	EventUsecase event.Usecase
}

func NewHttpEventHandler(e *echo.Echo, eu event.Usecase) {
	eh := &HttpEventHandler{eu}

	e.GET("/events/:id", eh.FetchEventByID)
	e.GET("/events", eh.FetchEvents)
	e.POST("/events", eh.StoreEvent)
}

func (eh *HttpEventHandler) StoreEvent(c echo.Context) error {
	e := new(models.Event)
	if err := c.Bind(e); err != nil {
		return err
	}

	if err := eh.EventUsecase.Store(e); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, e)
}

func (eh *HttpEventHandler) FetchEvents(c echo.Context) error {
	events, err := eh.EventUsecase.FetchAll()
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, events)
}

func (eh *HttpEventHandler) FetchEventByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	u, err := eh.EventUsecase.Fetch(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, u)
}
