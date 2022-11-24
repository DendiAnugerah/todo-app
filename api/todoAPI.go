package api

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/DendiAnugerah/Todo/model"
	"github.com/labstack/echo/v4"
)

func (A *API) addTodo(c echo.Context) error {
	username := fmt.Sprintf("%s", c.Request().Context().Value("username"))

	var todo model.Todo
	err := json.NewDecoder(c.Request().Body).Decode(&todo)
	if err != nil {
		return c.JSON(400, model.ErrorResponse{Error: err.Error()})
	}

	err = A.TodoRepo.AddTodo(todo)
	if err != nil {
		return c.JSON(500, model.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, model.SuccessResponse{Username: username, Message: fmt.Sprintf("Task %s added!", todo.Task)})
}

func (A *API) listTodo(c echo.Context) error {
	res, err := A.TodoRepo.ReadTodo()
	if err != nil {
		return c.JSON(500, model.ErrorResponse{Error: err.Error()})
	}

	if len(res) == 0 {
		return c.JSON(200, model.SuccessResponse{Message: "No task found!"})
	}

	c.JSON(200, res)
	return json.NewEncoder(c.Response()).Encode(res)
}

func (A *API) changeTodoStatus(c echo.Context) error {
	username := fmt.Sprintf("%s", c.Request().Context().Value("username"))

	var model model.ToggleTodoReq
	err := json.NewDecoder(c.Request().Body).Decode(&model)
	if err != nil {
		return c.JSON(400, nil)
	}

	err = A.TodoRepo.UpdateTodoStatus(model.ID, model.Done)
	if err != nil {
		return c.JSON(500, nil)
	}

	c.JSON(200, nil)
	return json.NewEncoder(c.Response()).Encode(username)
}

func (A *API) deleteTodo(c echo.Context) error {
	username := fmt.Sprintf("%s", c.Request().Context().Value("username"))

	id := c.Request().URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, model.ErrorResponse{Error: err.Error()})
	}

	err = A.TodoRepo.DeleteTodo(idInt)
	if err != nil {
		return c.JSON(500, model.ErrorResponse{Error: err.Error()})
	}

	c.JSON(200, nil)
	return json.NewEncoder(c.Response()).Encode(username)
}
