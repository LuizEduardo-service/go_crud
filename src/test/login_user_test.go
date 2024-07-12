package test

import (
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
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {

	t.Run("quando_relaiza_login_retorna_valido", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		recorderLoginUser := httptest.NewRecorder()
		ctxLoginUser := GetTestGinContext(recorderLoginUser)

		// CRIANDO USUARIO =======================
		email := fmt.Sprintf("%d@test.com", rand.Int())
		password := fmt.Sprintf("%d@@@$$##", rand.Int())

		userRequest := request.UserRequest{
			Email:    email,
			Password: password,
			Name:     "luiz",
			Age:      31,
		}

		b, _ := json.Marshal(userRequest)
		newReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctx, gin.Params{}, url.Values{}, "POST", newReader)
		UserController.CreateUser(ctx)

		// REALIZANDO LOGIN DO USUARIO =======================
		loginUserRequest := request.UserLogin{
			Email:    email,
			Password: userRequest.Password,
		}

		us, _ := json.Marshal(loginUserRequest)
		newUserReader := io.NopCloser(strings.NewReader(string(us)))
		MakeRequest(ctxLoginUser, gin.Params{}, url.Values{}, "POST", newUserReader)
		UserController.LoginUser(ctxLoginUser)

		assert.EqualValues(t, http.StatusOK, recorderLoginUser.Result().StatusCode)
		assert.NotEmpty(t, recorderLoginUser.Result().Header.Get("Authorization"))
	})

	t.Run("quando_realiza_login_retorna_invalido", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		recorderLoginUser := httptest.NewRecorder()
		ctxLoginUser := GetTestGinContext(recorderLoginUser)

		// CRIANDO USUARIO =======================
		email := fmt.Sprintf("%d@test.com", rand.Int())
		password := fmt.Sprintf("%d@@@$$##", rand.Int())

		userRequest := request.UserRequest{
			Email:    email,
			Password: password,
			Name:     "luiz",
			Age:      31,
		}

		b, _ := json.Marshal(userRequest)
		newReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctx, gin.Params{}, url.Values{}, "POST", newReader)
		UserController.CreateUser(ctx)

		// REALIZANDO LOGIN DO USUARIO =======================
		loginUserRequest := request.UserLogin{
			Email:    "erro@erro.com",
			Password: "@@@@@126464$$",
		}

		us, _ := json.Marshal(loginUserRequest)
		newUserReader := io.NopCloser(strings.NewReader(string(us)))
		MakeRequest(ctxLoginUser, gin.Params{}, url.Values{}, "POST", newUserReader)
		UserController.LoginUser(ctxLoginUser)

		assert.EqualValues(t, http.StatusForbidden, recorderLoginUser.Result().StatusCode)
		assert.Empty(t, recorderLoginUser.Result().Header.Get("Authorization"))
	})
}
