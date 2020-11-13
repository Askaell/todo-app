package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"

	"github.com/spf13/viper"

	"github.com/Askaell/todo-app/pkg/repository"
	"github.com/Askaell/todo-app/pkg/service"

	"github.com/Askaell/todo-app/pkg/handler"

	"github.com/Askaell/todo-app"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error inirializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repository := repository.NewRepository(db)
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
