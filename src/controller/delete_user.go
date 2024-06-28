package controller

import (
	"net/http"

	"github.com/LuizEduardo-service/go_crud/src/configuration/logger"
	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Iniciando exclusão de usuario", zap.String("journey", "deleteUser"))

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Valor incorreto")
		c.JSON(errRest.Code, errRest)
		return
	}

	err := uc.service.DeleteUser(userId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Usuario Deletado com sucesso", zap.String("journey", "deleteUser"))

	c.Status(http.StatusOK)
}
