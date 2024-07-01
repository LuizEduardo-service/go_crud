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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_LoginUser(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	service := mocks.NewMockUserDomainService(control)
	controller := NewControllerInterface(service)

	t.Run("quando_json_login_retorna_invalido", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserLogin{
			Email:    "luiz@teste.com",
			Password: "luiz@@@@@",
		}

		userDomain := model.NewUserLoginDomain(
			userRequest.Email,
			userRequest.Password,
		)

		b, _ := json.Marshal(userDomain)
		newReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", newReader)
		controller.LoginUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)

	})

	t.Run("quando_json_e_valido_e_service_retorna_invalido", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserLogin{
			Email:    "luiz@test.com",
			Password: "luiz@@@@@",
		}

		userDomain := model.NewUserLoginDomain(
			userRequest.Email,
			userRequest.Password,
		)

		b, _ := json.Marshal(userRequest)
		newReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().LoginUserServices(userDomain).Return(nil, "", rest_err.NewBadRequestError("Usuario Não Localizado!"))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", newReader)
		controller.LoginUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)

	})

	t.Run("quando_login_user_retorna_sucesso", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		mock_token := primitive.NewObjectID().Hex()

		userRequest := request.UserLogin{
			Email:    "luiz@test.com",
			Password: "luiz@@@@@",
		}

		userDomain := model.NewUserLoginDomain(
			userRequest.Email,
			userRequest.Password,
		)

		b, _ := json.Marshal(userRequest)
		newReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().LoginUserServices(userDomain).Return(userDomain, mock_token, nil)

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", newReader)
		controller.LoginUser(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, recorder.Header().Values("authorization")[0], mock_token)

	})
}
