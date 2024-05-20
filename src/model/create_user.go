package model

import (
	"fmt"

	"github.com/LuizEduardo-service/go_crud/src/configuration/logger"
	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Criando Model createUser", zap.String("journey", "createUser"))
	ud.EncryptPassword()
	fmt.Println((ud.Password))
	return nil
}
