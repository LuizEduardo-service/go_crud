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

// LoginUser allows a user to log in and obtain an authentication token.
// @Summary User Login
// @Description Allows a user to log in and receive an authentication token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param userLogin body request.UserLogin true "User login credentials"

// @Header 200 {string} Authorization "Authentication token"
// @Failure 403 {object} rest_err.RestErr "Error: Invalid login credentials"
// @Router /login [post]
func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Inciando login de Usuario", zap.String("journey", "userLogin"))

	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro validação usuario", err,
			zap.String("journey", "userLogin"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return

	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, token, err := uc.service.LoginUserServices(domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Usuario logado com sucesso", zap.String("journey", "userLogin"))
	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
