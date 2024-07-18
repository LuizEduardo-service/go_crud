package controller

import (
	"net/http"
	"net/mail"

	"github.com/LuizEduardo-service/go_crud/src/configuration/logger"
	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"github.com/LuizEduardo-service/go_crud/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// FindUserByEmail retrieves user information based on the provided email.
// @Summary Find User by Email
// @Description Retrieves user details based on the email provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userEmail path string true "Email of the user to be retrieved"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} response.UserResponse "User information retrieved successfully"
// @Failure 400 {object} rest_err.RestErr "Error: Invalid user ID"
// @Failure 404 {object} rest_err.RestErr "User not found"
// @Router /getUserByEmail/{userEmail} [get]
func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Iniciando Pesquisa por ID", zap.String("journey", "findUserByID"))

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

// FindUserByEmail retrieves user information based on the provided user Email.
// @Summary Find User by Email
// @Description Retrieves user details based on the user Email provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userEmail path string true "Email of the user to be retrieved"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} response.UserResponse "User information retrieved successfully"
// @Failure 400 {object} rest_err.RestErr "Error: Invalid user Email"
// @Failure 404 {object} rest_err.RestErr "User not found"
// @Router /getUserById/{userId} [get]
func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Iniciando Pesquisa por Email", zap.String("journey", "FindUserByEmail"))

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
