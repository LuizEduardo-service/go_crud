package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/LuizEduardo-service/go_crud/src/controller/model/request"
	"github.com/LuizEduardo-service/go_crud/src/model/repository/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateUser(t *testing.T) {

	t.Run("quando_criar_email_retorna_existente", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		email := fmt.Sprintf("%d@test.com", rand.Int())

		_, err := Database.Collection("users_test").InsertOne(context.Background(), bson.M{"name": t.Name(), "email": email})
		if err != nil {
			t.Fatal(err)
			return
		}
		// CIRANDO REQUEST PARA CONSULTA
		userRequest := request.UserRequest{
			Email:    email,
			Password: "123",
			Name:     "luiz",
			Age:      31,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctx, gin.Params{}, url.Values{}, "POST", stringReader)
		UserController.CreateUser(ctx)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("quando_usuario_nao_e_cadastrado_na_base", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		email := fmt.Sprintf("%d@test.com", rand.Int())

		_, err := Database.Collection("users_test").InsertOne(context.Background(), bson.M{"name": t.Name(), "email": email})
		if err != nil {
			t.Fatal(err)
			return
		}
		// CIRANDO REQUEST PARA CONSULTA
		userRequest := request.UserRequest{
			Email:    email,
			Password: "123",
			Name:     "luiz",
			Age:      31,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctx, gin.Params{}, url.Values{}, "POST", stringReader)
		UserController.CreateUser(ctx)

		// CRIAR ENTITY PARA DESCARREGAR O RESULTADO
		userEntity := entity.UserEntity{}
		filter := bson.D{{Key: "email", Value: email}}

		_ = Database.Collection("users_test").FindOne(context.Background(), filter).Decode(&userEntity)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, userEntity.Email, userRequest.Email)
		assert.EqualValues(t, userEntity.Password, userRequest.Password)
		assert.EqualValues(t, userEntity.Name, userRequest.Name)
		assert.EqualValues(t, userEntity.Age, userRequest.Age)
	})
}
