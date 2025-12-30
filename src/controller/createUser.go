package controller

import (
	"net/http"

	"github.com/Tommy-jeff/first_go_crud/src/configs/logger"
	"github.com/Tommy-jeff/first_go_crud/src/configs/validation"
	"github.com/Tommy-jeff/first_go_crud/src/controller/model/request"
	"github.com/Tommy-jeff/first_go_crud/src/model"
	"github.com/Tommy-jeff/first_go_crud/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

// Função que cria um user
func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser Controller",
		zap.String("Journey", "createUser"),
	)

	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("Error trying to validade user info", err,
			zap.String("Journey", "createUser"),
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	userDomain := model.NewUserDomain(
		UserRequest.Email,
		UserRequest.Password,
		UserRequest.Name,
		UserRequest.Age,
	)
	userService := service.NewUserDomainService()
	
	id, err := userService.CreateUser(userDomain)
	
	if err != nil {
		logger.Error("Error trying to create user", err,
			zap.String("Journey", "createUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("user created succesfully",
		zap.String("Journey", "createUser"),
	)

	c.JSON(http.StatusOK, id)
}
