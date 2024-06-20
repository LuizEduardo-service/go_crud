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

func TestUserDomainInterface_FindUserByIDServices(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repo := mocks.NewMockUserRepository(control)
	service := NewUserDomainService(repo)

	t.Run("quando_existe_usuario_retorna_sucesso", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "luiz", 31)
		userDomain.SetID(id)

		repo.EXPECT().FindUserByID(id).Return(userDomain, nil)
		userDomainReturn, err := service.FindUserByIDServices(id)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), id)
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())

	})

	t.Run("quando_nao_existe_usuario_retorna_erro", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repo.EXPECT().FindUserByID(id).Return(nil, rest_err.NewNotFoundError("Usuario não encontrado!"))
		userDomainReturn, err := service.FindUserByIDServices(id)

		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Usuario não encontrado!")

	})
}

func TestUserDomainInterface_FindUserByEmailServices(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repo := mocks.NewMockUserRepository(control)
	service := NewUserDomainService(repo)

	t.Run("quando_existe_usuario_retorna_Sucesso", func(t *testing.T) {
		// criar id
		id := primitive.NewObjectID().Hex()
		email := "teste@success.com"
		// criar domain
		userDomain := model.NewUserDomain(email, "123", "luiz", 31)
		// setar id
		userDomain.SetID(id)

		// mocar a resposta esperada
		repo.EXPECT().FindUserByEmail(email).Return(userDomain, nil)
		// realizar a consulta no service por email
		userDomainReturn, err := service.FindUserByEmailServices(email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), id)
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())

		// usar assert para verificar se tudo retornou com sucesso
	})

	t.Run("quando_nao_existe_usuario_retorna_erro", func(t *testing.T) {
		// criar email
		email := "teste@error.com"
		// mocar resposta
		repo.EXPECT().FindUserByEmail(email).Return(nil, rest_err.NewNotFoundError("Usuario não localizado por email"))
		// realizar consulta por email no service
		userDomainReturn, err := service.FindUserByEmailServices(email)
		// realizar asserts

		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Usuario não localizado por email")

	})
}

func TestUserDomainInterface_FindUserByEmailAndPasswordServices(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repo := mocks.NewMockUserRepository(control)
	service := &userDomainService{repo}

	t.Run("Quando_consulta_email_e_password_retorna_sucesso", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "teste@test.com"
		password := strconv.FormatInt(rand.Int63(), 10)

		userDomain := model.NewUserDomain(email, password, "luiz", 31)
		userDomain.SetID(id)

		repo.EXPECT().FindUserByEmailAndPassword(email, password).Return(userDomain, nil)

		userDomainReturn, err := service.findUserByEmailAndPasswordServices(email, password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())

	})

	t.Run("Quando_consulta_email_password_retorna_erro", func(t *testing.T) {
		email := "test@test.com"
		password := strconv.FormatInt(rand.Int63(), 10)

		repo.EXPECT().FindUserByEmailAndPassword(email, password).Return(nil, rest_err.NewNotFoundError("Usuario não localizado pesquisando por email e password"))

		userDomainReturn, err := service.findUserByEmailAndPasswordServices(email, password)

		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Usuario não localizado pesquisando por email e password")

	})

}
