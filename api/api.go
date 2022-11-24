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

	c.POST("/user/register", api.RegisterUser)
	c.POST("/user/login", api.Login)
	// c.GET("/user/logout", api.Logout)

	return api
}

func (A *API) Handler() *echo.Echo {
	return API
}
func (A *API) Start() {
	e := echo.New()

	fmt.Println("Starting server on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
