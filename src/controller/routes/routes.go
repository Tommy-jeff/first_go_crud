package routes

import (
	"github.com/Tommy-jeff/first_go_crud/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:id", controller.FindUserById)
	r.GET("/getUserByEmail/:email", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:id", controller.UpdateUser)
	r.DELETE("/deleteUser/:id", controller.DeleteUser)

}