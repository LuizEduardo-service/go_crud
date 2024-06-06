package service

import (
	"github.com/LuizEduardo-service/go_crud/src/configuration/logger"
	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"github.com/LuizEduardo-service/go_crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, string, *rest_err.RestErr) {

	userDomain.EncryptPassword()
	user, err := ud.findUserByEmailAndPasswordServices(
		userDomain.GetEmail(),
		userDomain.GetPassword())

	if err != nil {
		return nil, "", err
	}

	token, err := userDomain.GenerateToken()

	if err != nil {
		return nil, "", err
	}
	logger.Info("Usuario localizado", zap.String("journey", "loginUser"))

	return user, token, nil
}
