package main

import (
	"log"

	"github.com/spf13/viper"

	"github.com/Askaell/todo-app/pkg/repository"
	"github.com/Askaell/todo-app/pkg/service"

	"github.com/Askaell/todo-app/pkg/handler"

	"github.com/Askaell/todo-app"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("errpr inirializing configs: %s", err.Error())
	}

	repository := repository.NewRepository()
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server, %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
