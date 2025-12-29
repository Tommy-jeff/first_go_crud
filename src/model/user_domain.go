package model

import (
	"crypto/md5"
	"encoding/hex"

	resterr "github.com/Tommy-jeff/first_go_crud/src/configs/rest_err"
)

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

// Uma struct implementa uma interface automaticamente se ela possuir todos os métodos exigidos pela interface, com as mesmas assinaturas.
// Quando você implementar todos os métodos da interface, o compilador automaticamente considera que *UserDomain satisfaz UserDomainInterface.
func NewUserDomain(email string, password string, name string, age int8) UserDomainInterface {
	return &UserDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *resterr.RestErr
	UpdateUser(string) *resterr.RestErr
	FindUser(string) (*UserDomain, *resterr.RestErr)
	DeleteUser(string) *resterr.RestErr
}