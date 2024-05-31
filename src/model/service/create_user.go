package service

import (
	"github.com/LuizEduardo-service/go_crud/src/configuration/logger"
	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"github.com/LuizEduardo-service/go_crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Criando Model createUser", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)

	if err != nil {
		return nil, err
	}

	return userDomainRepository, nil
}
