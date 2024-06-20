package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_DeleteUser(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	// setar as variaveis de ambiente
	os.Setenv("DATABASE_COLLECTION", collection_name)
	defer os.Clearenv()

	//criar novo banco de test
	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("testando_exclusao_usuario", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		err := repo.DeleteUser("1231")

		assert.Nil(t, err)
	})

	mtestDb.Run("testando_retorno_de_erro", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		err := repo.DeleteUser("test")

		assert.NotNil(t, err)

	})
}
