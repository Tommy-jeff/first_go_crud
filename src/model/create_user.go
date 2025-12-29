package model

import (
	"github.com/Tommy-jeff/first_go_crud/src/configs/logger"
	resterr "github.com/Tommy-jeff/first_go_crud/src/configs/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *resterr.RestErr {
	logger.Info("Init CreateUser model",
		zap.String("Journey", "createUser"),
	)

	ud.EncryptPassword()

	// fmt.Println("Password encrypted: ",ud.Password)

	return nil
}
