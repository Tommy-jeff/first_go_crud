package service

import (
	resterr "github.com/Tommy-jeff/first_go_crud/src/configs/rest_err"
	"github.com/Tommy-jeff/first_go_crud/src/model"
)

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (int, *resterr.RestErr)
	UpdateUser(string, model.UserDomainInterface) *resterr.RestErr
	FindUser(string) (*model.UserDomainInterface, *resterr.RestErr)
	DeleteUser(string) *resterr.RestErr
}

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}