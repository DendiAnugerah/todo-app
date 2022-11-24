package main

import (
	"github.com/DendiAnugerah/Todo/api"
	"github.com/DendiAnugerah/Todo/config"
	"github.com/DendiAnugerah/Todo/model"
	"github.com/DendiAnugerah/Todo/repository"
)

func main() {
	db := config.NewDB()

	conn := db.Connection()
	conn.AutoMigrate(&model.User{}, &model.Todo{}, &model.Session{})

	usersRepo := repository.NewUserRepository(conn)
	todosRepo := repository.NewTodoRepository(conn)
	sessionsRepo := repository.NewSessionRepository(conn)

	API := api.NewAPI(todosRepo, usersRepo, sessionsRepo)
	API.Start()
}
