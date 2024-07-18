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

// var (
// 	UserDomainInterface model.UserDomainInterface
// )

// CreateUser Creates a new user
// @Summary Create a new user
// @Description Create a new user with the provided user information
// @Tags Users
// @Accept json
// @Produce json
// @Param userRequest body request.UserRequest true "User information for registration"
// Success 200 {object} response.UserResponse
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /createUser [post]
func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Inciando Criação de Usuario", zap.String("journey", "createUser"))

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
		return
	}

	logger.Info("Usuario Criado com sucesso", zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
