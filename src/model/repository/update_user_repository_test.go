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

func TestUserRepository_UpdateUser(t *testing.T) {

	database_name := "user_database_test"
	collection_name := "user_collection_test"

	os.Setenv("DATABASE_COLLECTION", collection_name)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("testando_atulizar_usuario", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		userDomain := model.NewUserDomain(
			"test@test2.com",
			"teste",
			"luiz",
			30,
		)
		userDomain.SetID(primitive.NewObjectID().Hex())
		err := repo.UpdateUser(userDomain.GetID(), userDomain)

		assert.Nil(t, err)

		mtestDb.Run("testando_retorno_de_erro_update", func(mt *mtest.T) {
			mt.AddMockResponses(bson.D{
				{Key: "ok", Value: 0},
			})

			databaseMock := mt.Client.Database(database_name)
			repo := NewUserRepository(databaseMock)
			userDomain := model.NewUserDomain(
				"test@test2.com",
				"teste",
				"luiz",
				30,
			)
			userDomain.SetID(primitive.NewObjectID().Hex())
			err := repo.UpdateUser(userDomain.GetID(), userDomain)

			assert.NotNil(t, err)

		})

	})
}
