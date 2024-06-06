package controller

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/LuizEduardo-service/go_crud/src/configuration/logger"
	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"github.com/LuizEduardo-service/go_crud/src/model"
	"github.com/LuizEduardo-service/go_crud/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Iniciando Pesquisa por ID", zap.String("journey", "findUserByID"))

	user, err := model.VerifyToken(c.Request.Header.Get("authorization"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("Usuario autenticado: %#v", user))

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := rest_err.NewBadRequestError(
			"ID não é valido!",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Iniciando Pesquisa por Email", zap.String("journey", "FindUserByEmail"))

	user, err := model.VerifyToken(c.Request.Header.Get("authorization"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("Usuario autenticado: %#v", user))
	email := c.Param("userEmail")

	if _, err := mail.ParseAddress(email); err != nil {
		errorMessage := rest_err.NewBadRequestError(
			"ID não é valido!",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(email)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))

}
