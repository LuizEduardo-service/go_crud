package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDeleteUser(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx := GetTestGinContext(recorder)
	id := primitive.NewObjectID()

	// inserindo usuario
	_, err := Database.Collection("users_test").InsertOne(context.Background(), bson.M{"_id": id, "name": t.Name(), "email": "test@test.com"})
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
	MakeRequest(ctx, param, url.Values{}, "DELETE", nil)
	// chamar a função do controller global
	UserController.DeleteUser(ctx)

	// assert retornando estatus erro
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	filter := bson.D{{Key: "_id", Value: id}}
	result := Database.Collection("users_test").FindOne(context.Background(), filter)

	assert.NotNil(t, result.Err())

}
