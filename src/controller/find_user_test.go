package controller

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"github.com/LuizEduardo-service/go_crud/src/model"
	"github.com/LuizEduardo-service/go_crud/src/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_FindUserByEmail(t *testing.T) {

	control := gomock.NewController(t)
	defer control.Finish()

	service := mocks.NewMockUserDomainService(control)

	controller := NewControllerInterface(service)

	t.Run("quando_retorna_email_invalido", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "TEST_ERROR",
			},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("quando_retorna_erro_service", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "teste@test.com",
			},
		}

		service.EXPECT().FindUserByEmailServices("teste@test.com").Return(nil, rest_err.NewInternalServerError("Erro ao localizar usuario por email!"))

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("quando_service_retorna_sucesso", func(t *testing.T) {
		recoder := httptest.NewRecorder()
		context := GetTestGinContext(recoder)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "test@test.com",
			},
		}

		userDomain := model.NewUserDomain("test@test.com", "123", "luiz", 31)
		service.EXPECT().FindUserByEmailServices(userDomain.GetEmail()).Return(userDomain, nil)
		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusOK, recoder.Code)
	})
}

func TestUserControllerInterface_FindUserById(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	service := mocks.NewMockUserDomainService(control)

	controller := NewControllerInterface(service)

	t.Run("quando_retorna_id_invalido", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userId",
				Value: "test",
			},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByID(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("quando_retorna_erro_service", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().FindUserByIDServices(id).Return(nil, rest_err.NewInternalServerError("Erro ao localizar usuario pelo ID!"))

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByID(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)

	})

	t.Run("quando_retorna_sucesso_id", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}
		userDomain := model.NewUserDomain("teste@test.com", "123", "luiz", 31)
		service.EXPECT().FindUserByIDServices(id).Return(userDomain, nil)

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByID(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)

	})
}

func GetTestGinContext(recoder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recoder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}
func MakeRequest(
	c *gin.Context,
	param gin.Params,
	u url.Values,
	method string,
	body io.ReadCloser,
) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param
	c.Request.Body = body
	c.Request.URL.RawQuery = u.Encode()
}
