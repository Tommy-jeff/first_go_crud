package routes

import (
	"github.com/Tommy-jeff/first_go_crud/src/controller"
	"github.com/gin-gonic/gin"
)


/// Função que inicializa os nossos endpoints e as funções as quais são chamadas para cada rota.
/// Importante: cada função recebe um contexto do gin, algo que é não neccessário ser declado efetivamente porque esse contexto já é propagado quando um requisição é recebida.
func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:id", controller.FindUserById)
	r.GET("/getUserByEmail/:email", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:id", controller.UpdateUser)
	r.DELETE("/deleteUser/:id", controller.DeleteUser)

}