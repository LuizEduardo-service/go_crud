package repository

import (
	"context"
	"os"

	"github.com/LuizEduardo-service/go_crud/src/configuration/logger"
	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"github.com/LuizEduardo-service/go_crud/src/model"
	"github.com/LuizEduardo-service/go_crud/src/model/repository/entity"
	"github.com/LuizEduardo-service/go_crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmailAndPassword(
	email, password string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Iniciando Busca de usuario no repositorio login",
		zap.String("journey", "findUserByEmail"))

	collection_name := os.Getenv(DATABASE_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "Usuario não localizado! Verifique usuario e senha:"

			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Erro ao tentar localizar usuario"
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	return converter.ConvertEntityToDomain(*userEntity), nil

}
