package controller

import (
	"net/http"

	"github.com/LuizEduardo-service/go_crud/src/configuration/logger"
	"github.com/LuizEduardo-service/go_crud/src/configuration/validation"
	"github.com/LuizEduardo-service/go_crud/src/controller/model/request"
	"github.com/LuizEduardo-service/go_crud/src/model"
	"github.com/LuizEduardo-service/go_crud/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Iniando Criação de Usuario", zap.String("journey", "createUser"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro validação usuario", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return

	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domainResult, err := uc.service.CreateUserServices(domain)
	if err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("Usuario Criado com sucesso", zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
