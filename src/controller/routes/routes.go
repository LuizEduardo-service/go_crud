package routes

import (
	"github.com/LuizEduardo-service/go_crud/src/controller"
	"github.com/LuizEduardo-service/go_crud/src/model"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserId/:userId", model.VerifyTokenMiddleware, userController.FindUserByID)
	r.GET("/getUserEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
	r.POST("/login", userController.LoginUser)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
