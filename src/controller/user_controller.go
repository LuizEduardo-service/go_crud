package controller

import (
	"github.com/LuizEduardo-service/go_crud/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	UpdateUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
