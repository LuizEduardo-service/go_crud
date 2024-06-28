package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"github.com/LuizEduardo-service/go_crud/src/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_DeleteUserById(t *testing.T) {

	control := gomock.NewController(t)
	defer control.Finish()

	service := mocks.NewMockUserDomainService(control)
	controller := NewControllerInterface(service)

	t.Run("quando_id_invalido_retorna_erro", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userId",
				Value: "teste",
			},
		}

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("quando_id_valido_e_service_retorna_erro", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().DeleteUser(id).Return(rest_err.NewInternalServerError("Erro ao excluir Usuario!"))

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("quando_id_valido_retorna_Sucesso", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().DeleteUser(id).Return(nil)
		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}
