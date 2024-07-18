package controller

import (
	"net/http"

	"github.com/LuizEduardo-service/go_crud/src/configuration/logger"
	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"github.com/LuizEduardo-service/go_crud/src/configuration/validation"
	"github.com/LuizEduardo-service/go_crud/src/controller/model/request"
	"github.com/LuizEduardo-service/go_crud/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// UpdateUser updates user information with the specified ID.
// @Summary Update User
// @Description Updates user details based on the ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be updated"
// @Param userRequest body request.UserUpdateRequest true "User information for update"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /updateUser/{userId} [put]
func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Iniciando alteração de usuario", zap.String("journey", "updateUser"))

	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro validação usuario", err,
			zap.String("journey", "updateUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return

	}

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Valor incorreto")
		c.JSON(errRest.Code, errRest)
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("Usuario Criado com sucesso", zap.String("journey", "updateUser"))

	c.Status(http.StatusOK)
}
