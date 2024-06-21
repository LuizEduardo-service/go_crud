package service

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"github.com/LuizEduardo-service/go_crud/src/model"
	"github.com/LuizEduardo-service/go_crud/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainInterface_DeleteUserServices(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repo := mocks.NewMockUserRepository(control)
	service := NewUserDomainService(repo)

	t.Run("quando_exclui_usuario_retorna_Sucesso", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@test.com"
		password := strconv.FormatInt(rand.Int63(), 10)
		name := "luiz"
		age := 31

		userDomain := model.NewUserDomain(email, password, name, int8(age))
		userDomain.SetID(id)

		repo.EXPECT().DeleteUser(id).Return(nil)
		err := service.DeleteUser(id)

		assert.Nil(t, err)

	})

	t.Run("quando_exclui_usuario_retorna_erro", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@test.com"
		password := strconv.FormatInt(rand.Int63(), 10)
		name := "luiz"
		age := 31

		userDomain := model.NewUserDomain(email, password, name, int8(age))
		userDomain.SetID(id)

		repo.EXPECT().DeleteUser(id).Return(rest_err.NewInternalServerError("Erro ao excluir usuario"))

		err := service.DeleteUser(id)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Erro ao excluir usuario")
	})
}
