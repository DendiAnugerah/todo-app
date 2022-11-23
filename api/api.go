package api

import (
	"fmt"

	"github.com/DendiAnugerah/Todo/repository"
	"github.com/labstack/echo/v4"
)

type API struct {
	TodoRepo    repository.TodoRepository
	UserRepo    repository.UserRepository
	SessionRepo repository.SessionRepository
}

func NewAPI(todoRepo repository.TodoRepository, userRepo repository.UserRepository, sessionRepo repository.SessionRepository) API {
	api := API{todoRepo, userRepo, sessionRepo}
	c := echo.New()

	c.POST("/register", api.RegisterUser)
	c.POST("/login", api.Login)

	return api
}

func (A *API) Start() {
	e := echo.New()

	fmt.Println("Starting server on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
