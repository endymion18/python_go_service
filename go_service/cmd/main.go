package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go_service"
	"go_service/pkg/handler"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init configs %s", err.Error())
	}

	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("error loading env %s", err.Error())
	}

	handlers := new(handler.Handler)
	server := new(go_service.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
