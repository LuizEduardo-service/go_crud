package test

import (
	"context"
	"encoding/json"
	"io"
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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUpdateUser(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx := GetTestGinContext(recorder)
	id := primitive.NewObjectID()

	_, err := Database.Collection("users").InsertOne(context.Background(), bson.M{"_id": id, "name": "luiz", "age": 10, "email": "test@test.com"})
	if err != nil {
		t.Fatal(err)
		return
	}

	param := []gin.Param{
		{
			Key:   "userId",
			Value: id.Hex(),
		},
	}

	userRequest := request.UserUpdateRequest{
		Name: "Luiz Eduardo",
		Age:  32,
	}

	b, _ := json.Marshal(userRequest)
	newReader := io.NopCloser(strings.NewReader(string(b)))

	MakeRequest(ctx, param, url.Values{}, "PUT", newReader)
	UserController.UpdateUser(ctx)

	userEntity := entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: id}}
	_ = Database.Collection("users").FindOne(context.Background(), filter).Decode(&userEntity)

	assert.EqualValues(t, http.StatusOK, recorder.Result().StatusCode)
	assert.EqualValues(t, userEntity.Name, userRequest.Name)
	assert.EqualValues(t, userEntity.Age, userRequest.Age)
}
