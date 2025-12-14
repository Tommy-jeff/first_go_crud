package controller

import (
	"github.com/Tommy-jeff/first_go_crud/src/configs/validation"
	"github.com/Tommy-jeff/first_go_crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

/// Função que cria um user
func CreateUser(c *gin.Context) {

	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
	}

}