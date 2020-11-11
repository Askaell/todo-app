package main

import (
	"log"

	"github.com/Askaell/todo-app/pkg/repository"
	"github.com/Askaell/todo-app/pkg/service"

	"github.com/Askaell/todo-app/pkg/handler"

	"github.com/Askaell/todo-app"
)

func main() {
	repository := repository.NewRepository()
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server, %s", err.Error())
	}
}
