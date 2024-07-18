package main

import (
	"context"
	"log"

	_ "github.com/LuizEduardo-service/go_crud/docs"
	"github.com/LuizEduardo-service/go_crud/src/configuration/database/mongodb"
	"github.com/LuizEduardo-service/go_crud/src/configuration/logger"
	"github.com/LuizEduardo-service/go_crud/src/controller"
	"github.com/LuizEduardo-service/go_crud/src/controller/routes"
	"github.com/LuizEduardo-service/go_crud/src/model/repository"
	"github.com/LuizEduardo-service/go_crud/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewControllerInterface(service)
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @host      localhost:8080
// @BasePath  /api/v1
// @schemes http
// @license MIT
func main() {
	logger.Info("Iniciando o Sistema")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Erro em tentar conectar ao database=%s \n", err.Error())
		return

	}

	userController := initDependencies(database)
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}

}
