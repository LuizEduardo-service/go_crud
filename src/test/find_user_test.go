package test

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/LuizEduardo-service/go_crud/src/controller"
	"github.com/LuizEduardo-service/go_crud/src/model/repository"
	"github.com/LuizEduardo-service/go_crud/src/model/service"
	"github.com/LuizEduardo-service/go_crud/src/test/connection"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// variavel global da interface de controller
var (
	UserController controller.UserControllerInterface
	Database       *mongo.Database
)

// criar testMain inicializador---------------------------
func TestMain(m *testing.M) {
	// setar variaveis de ambiente
	err := os.Setenv("DATABASE_NAME", "users_test")
	if err != nil {
		log.Fatalf("Erro ao setar variaveis de ambiente")
		return
	}
	// criar conexão
	closeConnection := func() {}
	Database, closeConnection := connection.OpenConnection()
	// criar repositorio
	repo := repository.NewUserRepository(Database)
	// criar service
	userService := service.NewUserDomainService(repo)
	// criar controller *global
	UserController = controller.NewControllerInterface(userService)
	// criar função anonima para o defer junto com o close connection
	defer func() {
		os.Clearenv()
		closeConnection()

	}()

	// ao criar testMain ja inserir os.exit
	os.Exit(m.Run())
}

func TestFindUserByEmail(t *testing.T) {
	Database, _ := connection.OpenConnection()
	t.Run("quando_email_retorna_invalido", func(t *testing.T) {
		// criar recorder
		recorder := httptest.NewRecorder()
		// criar contexto
		context := GetTestGinContext(recorder)
		// criar parametros usando array de gin
		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "test@test.com",
			},
		}
		// criar request usando a função Makerequest
		MakeRequest(context, param, url.Values{}, "GET", nil)
		// chamar a função do controller global
		UserController.FindUserByEmail(context)

		// assert retornando estatus erro
		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)

	})

	t.Run("quando_email_retorna_sucesso", func(t *testing.T) {
		// criar recorder
		recorder := httptest.NewRecorder()
		// criar contexto
		ctx := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		// inserindo usuario
		_, err := Database.
			Collection("users_test").
			InsertOne(context.Background(), bson.M{"_id": id, "name": t.Name(), "email": "test@test.com"})
		if err != nil {
			t.Fatal(err)
			return
		}
		// criar parametros usando array de gin
		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "test@test.com",
			},
		}
		// criar request usando a função Makerequest
		MakeRequest(ctx, param, url.Values{}, "GET", nil)
		// chamar a função do controller global
		UserController.FindUserByEmail(ctx)

		// assert retornando estatus erro
		assert.EqualValues(t, http.StatusOK, recorder.Code)

	})

}

func TestFindUserById(t *testing.T) {

	t.Run("quando_id_retorna_invalido", func(t *testing.T) {
		// criar recorder
		recorder := httptest.NewRecorder()
		// criar contexto
		context := GetTestGinContext(recorder)
		// criar parametros usando array de gin
		param := []gin.Param{
			{
				Key:   "userId",
				Value: "test",
			},
		}
		// criar request usando a função Makerequest
		MakeRequest(context, param, url.Values{}, "GET", nil)
		// chamar a função do controller global
		UserController.FindUserByID(context)

		// assert retornando estatus erro
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)

	})

	t.Run("quando_id_retorna_sucesso", func(t *testing.T) {
		// criar recorder
		recorder := httptest.NewRecorder()
		// criar contexto
		ctx := GetTestGinContext(recorder)
		id := primitive.NewObjectID()

		// inserindo usuario
		_, err := Database.
			Collection("users_test").
			InsertOne(context.Background(), bson.M{"_id": id, "name": t.Name(), "email": "test@test.com"})
		if err != nil {
			t.Fatal(err)
			return
		}
		// criar parametros usando array de gin
		param := []gin.Param{
			{
				Key:   "userId",
				Value: id.Hex(),
			},
		}
		// criar request usando a função Makerequest
		MakeRequest(ctx, param, url.Values{}, "GET", nil)
		// chamar a função do controller global
		UserController.FindUserByID(ctx)

		// assert retornando estatus erro
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
