package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/LuizEduardo-service/go_crud/src/configuration/logger"
	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"github.com/LuizEduardo-service/go_crud/src/model"
	"github.com/LuizEduardo-service/go_crud/src/model/repository/entity"
	"github.com/LuizEduardo-service/go_crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Iniciando Busca de usuario no repositorio",
		zap.String("journey", "findUserByEmail"))

	collection_name := os.Getenv(DATABASE_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"Usuario não localizado por email: %s", email,
			)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Erro ao tentar localizar usuario"
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	return converter.ConvertEntityToDomain(*userEntity), nil

}

func (ur *userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Iniciando Busca de usuario no repositorio",
		zap.String("journey", "findUserByid"))

	collection_name := os.Getenv(DATABASE_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectID}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"Usuario não localizado por id: %s", id,
			)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Erro ao tentar localizar usuario"
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	return converter.ConvertEntityToDomain(*userEntity), nil

}
