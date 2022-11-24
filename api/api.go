package api

import (
	"fmt"
	"net/http"

	"github.com/DendiAnugerah/Todo/repository"
)

type API struct {
	TodoRepo    repository.TodoRepository
	UserRepo    repository.UserRepository
	SessionRepo repository.SessionRepository
	mux         *http.ServeMux
}

func NewAPI(todoRepo repository.TodoRepository, userRepo repository.UserRepository, sessionRepo repository.SessionRepository) API {
	mux := http.NewServeMux()
	api := API{
		todoRepo,
		userRepo,
		sessionRepo,
		mux,
	}

	mux.Handle("/user/register", api.Post(http.HandlerFunc(api.Register)))
	mux.Handle("/user/login", api.Post(http.HandlerFunc(api.Login)))
	mux.Handle("/user/logout", api.Get(api.Auth(http.HandlerFunc(api.Logout))))

	mux.Handle("/todo/add", api.Post(api.Auth(http.HandlerFunc(api.addTodo))))
	mux.Handle("/todo/remove", api.Delete(api.Auth(http.HandlerFunc(api.deleteTodo))))
	mux.Handle("/todo/change-status", api.Put(api.Auth(http.HandlerFunc(api.changeTodoStatus))))
	mux.Handle("/todo/list", api.Get(api.Auth(http.HandlerFunc(api.listTodo))))

	return api
}

func (A *API) Handler() *http.ServeMux {
	return A.mux
}

func (A *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", A.Handler())
}
