package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/LuizEduardo-service/go_crud/src/model/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {

	database_name := "user_database_test"
	collection_name := "user_collection_test"

	os.Setenv("DATABASE_COLLECTION", collection_name)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("testando_pesquisa_por_email", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "test",
			Name:     "test",
			Age:      31,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch, convertEntityToBson(userEntity)))

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail(userEntity.Email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
	})

	mtestDb.Run("testando_erro_ao_pesquisar_por_email", func(mt *mtest.T) {

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch))

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("test")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)

	})

	mtestDb.Run("testando_documento_nao_encontrado", func(mt *mtest.T) {

		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch))

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("test")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)

	})

}

func convertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.ID},
		{Key: "email", Value: userEntity.Email},
		{Key: "password", Value: userEntity.Password},
		{Key: "name", Value: userEntity.Name},
		{Key: "age", Value: userEntity.Age},
	}
}
