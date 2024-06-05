package service

import (
	"github.com/LuizEduardo-service/go_crud/src/configuration/logger"
	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Excluindo", zap.String("journey", "deleteUser"))

	err := ud.userRepository.DeleteUser(userId)

	if err != nil {
		return err
	}

	return nil
}
