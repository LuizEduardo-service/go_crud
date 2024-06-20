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

func TestUserDomainInterface_CreateUserService(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repo := mocks.NewMockUserRepository(control)
	service := NewUserDomainService(repo)

	t.Run("Testando_criando_usuario_existente_retorna_erro", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "teste@test.com"
		password := strconv.FormatInt(rand.Int63(), 10)
		name := "luiz"
		age := 31
		userDomain := model.NewUserDomain(email, password, name, int8(age))
		userDomain.SetID(id)

		repo.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(userDomain, nil)

		user, err := service.CreateUserServices(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Email já Existe!")

	})

	t.Run("quando_nao_existe_usuario_retorna_erro", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "teste@test.com"
		password := strconv.FormatInt(rand.Int63(), 10)
		name := "luiz"
		age := 31

		userDomain := model.NewUserDomain(email, password, name, int8(age))
		userDomain.SetID(id)

		repo.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)
		repo.EXPECT().CreateUser(userDomain).Return(nil, rest_err.NewInternalServerError("Erro ao tentar criar usuario!"))
		user, err := service.CreateUserServices(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Erro ao tentar criar usuario!")
	})

	t.Run("quando_criacao_usuario_retorna_sucesso", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "teste@test.com"
		password := strconv.FormatInt(rand.Int63(), 10)
		name := "luiz"
		age := 31

		userDomain := model.NewUserDomain(email, password, name, int8(age))
		userDomain.SetID(id)

		repo.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)

		repo.EXPECT().CreateUser(userDomain).Return(userDomain, nil)

		user, err := service.CreateUserServices(userDomain)

		assert.Nil(t, err)
		assert.EqualValues(t, user.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, user.GetID(), userDomain.GetID())
		assert.EqualValues(t, user.GetName(), userDomain.GetName())
		assert.EqualValues(t, user.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, user.GetAge(), userDomain.GetAge())
	})
}
