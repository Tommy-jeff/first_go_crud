package service

import (
	"fmt"

	"github.com/Tommy-jeff/first_go_crud/src/configs/logger"
	resterr "github.com/Tommy-jeff/first_go_crud/src/configs/rest_err"
	"github.com/Tommy-jeff/first_go_crud/src/model"
	"go.uber.org/zap"
)

func (us *userDomainService) CreateUser(ui model.UserDomainInterface) (int, *resterr.RestErr) {
	logger.Info("Init CreateUser model",
		zap.String("Journey", "createUser"),
	)

	ui.EncryptPassword()

	fmt.Println("Password encrypted: ", ui.GetPassword())

	return 1, nil
}
