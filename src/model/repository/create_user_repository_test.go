package repository

import (
	"os"
	"testing"

	"github.com/LuizEduardo-service/go_crud/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_CreateUser(t *testing.T) {

	database_name := "user_database_test"
	collection_name := "user_collection_test"

	os.Setenv("DATABASE_COLLECTION", collection_name)
	defer os.Clearenv() //limpa as variaveis da memoria

	//mock de dados
	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.Close()

	mtestDb.Run("testando_criacao_usuario", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "Ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.CreateUser(model.NewUserDomain(
			"teste@gmail.com", "123456test", "luiz eduardo", 31,
		))
		_, errId := primitive.ObjectIDFromHex(userDomain.GetID())

		assert.Nil(t, err)
		assert.Nil(t, errId)
		assert.EqualValues(t, userDomain.GetEmail(), "teste@gmail.com")

	})
}
