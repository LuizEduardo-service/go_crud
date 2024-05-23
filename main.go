package main

import (
	"log"

	"github.com/LuizEduardo-service/go_crud/src/configuration/logger"
	"github.com/LuizEduardo-service/go_crud/src/controller"
	"github.com/LuizEduardo-service/go_crud/src/controller/routes"
	"github.com/LuizEduardo-service/go_crud/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Iniciando o Sistema")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// mongodb.InitConnection()
	// iniciando as dependencias
	service := service.NewUserDomainService()
	userController := controller.NewControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
