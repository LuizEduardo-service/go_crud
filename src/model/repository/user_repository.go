package repository

import (
	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"github.com/LuizEduardo-service/go_crud/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DATABASE_COLLECTION = "DATABASE_COLLECTION"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(
		userId string, userDomain model.UserDomainInterface,
	) *rest_err.RestErr
	DeleteUser(
		userId string,
	) *rest_err.RestErr

	FindUserByEmail(

		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByID(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmailAndPassword(
		email, password string,
	) (model.UserDomainInterface, *rest_err.RestErr)
}
