package controller

import (
	"fmt"

	"github.com/Tommy-jeff/first_go_crud/src/configs/rest_err"
	"github.com/Tommy-jeff/first_go_crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		restErr := resterr.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%v", err.Error()))
		c.JSON(restErr.Code, restErr)
	}

	fmt.Println(UserRequest)

}