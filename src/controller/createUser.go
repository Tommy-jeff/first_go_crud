package controller

import (
	"net/http"

	"github.com/Tommy-jeff/first_go_crud/src/configs/logger"
	"github.com/Tommy-jeff/first_go_crud/src/configs/validation"
	"github.com/Tommy-jeff/first_go_crud/src/controller/model/request"
	"github.com/Tommy-jeff/first_go_crud/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

// / Função que cria um user
func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser Controller",
		zapcore.Field{
			Key:    "journey",
			String: "createUser",
		},
	)

	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("Error trying to validade user info", err,
			zapcore.Field{
				Key:    "journey",
				String: "createUser",
			},
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	response := response.UserResponse{
		ID:    1,
		Name:  UserRequest.Name,
		Email: UserRequest.Email,
		Age:   UserRequest.Age,
	}

	logger.Info("user created succesfully",
		zapcore.Field{
			Key:    "journey",
			String: "createUser",
		},
	)

	c.JSON(http.StatusOK, response)
}
