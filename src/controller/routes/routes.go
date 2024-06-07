package routes

import (
	"github.com/LuizEduardo-service/go_crud/src/controller"
	"github.com/LuizEduardo-service/go_crud/src/model"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserId/:userId", model.VerifyTokenMiddleware, userController.FindUserByID)
	r.GET("/getUserEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
	r.POST("/login", userController.LoginUser)

}
