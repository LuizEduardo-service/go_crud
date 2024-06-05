package repository

import (
	"context"
	"os"

	"github.com/LuizEduardo-service/go_crud/src/configuration/logger"
	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *userRepository) DeleteUser(
	userId string,
) *rest_err.RestErr {
	logger.Info("Iniciando Exclusão de usuario no repositorio")
	collection_name := os.Getenv(DATABASE_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return rest_err.NewBadRequestError("erro na conversão do ID!")
	}

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}
