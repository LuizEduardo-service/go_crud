package service

import (
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"github.com/LuizEduardo-service/go_crud/src/model"
	"github.com/LuizEduardo-service/go_crud/src/test/mocks"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserDomainInterface_LoginUserServices(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repo := mocks.NewMockUserRepository(control)

	service := &userDomainService{repo}

	t.Run("quando_consulta_email_password_retorna erro", func(t *testing.T) {
		id := strconv.FormatInt(rand.Int63(), 10)
		email := "test@test.com"
		password := "1234"
		name := "luiz"
		age := 31

		userDomain := model.NewUserDomain(email, password, name, int8(age))
		userDomain.SetID(id)

		userDomainMock := model.NewUserDomain(
			userDomain.GetEmail(),
			userDomain.GetPassword(),
			userDomain.GetName(),
			userDomain.GetAge())

		userDomainMock.EncryptPassword()

		repo.EXPECT().FindUserByEmailAndPassword(
			userDomain.GetEmail(), userDomainMock.GetPassword()).Return(
			nil, rest_err.NewInternalServerError("Erro ao localizar usuario!"))

		user, token, err := service.LoginUserServices(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Empty(t, token)
		assert.EqualValues(t, err.Message, "Erro ao localizar usuario!")
	})

	t.Run("quando_gerar_token_retorna_erro", func(t *testing.T) {

		userDomainMock := mocks.NewMockUserDomainInterface(control)
		userDomainMock.EXPECT().GetEmail().Return("test@test.com")
		userDomainMock.EXPECT().GetPassword().Return("1234")
		userDomainMock.EXPECT().EncryptPassword()

		userDomainMock.EXPECT().GenerateToken().Return("", rest_err.NewInternalServerError("Erro ao gerar Token"))

		repo.EXPECT().FindUserByEmailAndPassword("test@test.com", "1234").Return(userDomainMock, nil)

		user, token, err := service.LoginUserServices(userDomainMock)

		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Erro ao gerar Token")
	})

	t.Run("quando_login_usuario_retorna_sucesso", func(t *testing.T) {
		id := strconv.FormatInt(rand.Int63(), 10)
		email := "test@test.com"
		password := "1234"
		name := "luiz"
		age := 31
		secret := "test"
		os.Setenv("JWT_SECRET_KEY", "test")
		defer os.Clearenv()
		userDomain := model.NewUserDomain(email, password, name, int8(age))
		userDomain.SetID(id)

		repo.EXPECT().FindUserByEmailAndPassword("test@test.com", gomock.Any()).Return(userDomain, nil)

		user, token, err := service.LoginUserServices(userDomain)

		assert.Nil(t, err)
		assert.EqualValues(t, user.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, user.GetID(), userDomain.GetID())
		assert.EqualValues(t, user.GetName(), userDomain.GetName())
		assert.EqualValues(t, user.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, user.GetAge(), userDomain.GetAge())

		tokenReturn, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
				return []byte(secret), nil
			}

			return nil, rest_err.NewBadRequestError("token Invalido")
		})

		_, ok := tokenReturn.Claims.(jwt.MapClaims)
		if !ok || !tokenReturn.Valid {
			t.FailNow()
			return
		}

	})
}
