package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"github.com/LuizEduardo-service/go_crud/src/controller/model/request"
	"github.com/LuizEduardo-service/go_crud/src/model"
	"github.com/LuizEduardo-service/go_crud/src/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_CreateUser(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	service := mocks.NewMockUserDomainService(control)
	controller := NewControllerInterface(service)

	t.Run("quando_json_retorna_invalido", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "Erro_email",
			Password: "1234",
			Name:     "luiz",
			Age:      31,
		}

		// conversão do userrequest em io.reader
		b, _ := json.Marshal(userRequest)
		reader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", reader)
		controller.CreateUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("quando_json_valido_e_retorna_erro _service", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "luiz@test.com",
			Password: "1234@@@@@@@",
			Name:     "luiz",
			Age:      31,
		}

		userDomain := model.NewUserDomain(
			userRequest.Email,
			userRequest.Password,
			userRequest.Name,
			userRequest.Age,
		)

		b, _ := json.Marshal(userRequest)
		newReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().CreateUserServices(userDomain).Return(nil, rest_err.NewInternalServerError("Erro ao criar Usuario"))
		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", newReader)
		controller.CreateUser(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("quando_criar_usuario_retorna_sucesso", func(t *testing.T) {
		recoder := httptest.NewRecorder()
		context := GetTestGinContext(recoder)

		userRequest := request.UserRequest{
			Email:    "luiz@test.com",
			Password: "1234@@@@@@@",
			Name:     "luiz",
			Age:      31,
		}

		userDomain := model.NewUserDomain(
			userRequest.Email,
			userRequest.Password,
			userRequest.Name,
			userRequest.Age,
		)

		b, _ := json.Marshal(userRequest)
		newReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().CreateUserServices(userDomain).Return(userDomain, nil)
		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", newReader)
		controller.CreateUser(context)

		assert.EqualValues(t, http.StatusOK, recoder.Code)
	})
}
