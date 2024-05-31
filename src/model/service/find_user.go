package service

import (
	"github.com/LuizEduardo-service/go_crud/src/configuration/logger"
	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"github.com/LuizEduardo-service/go_crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Iniciando Pesquisa por ID", zap.String("journey", "findUserID"))
	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Iniciando Pesquisa por email", zap.String("journey", "findUserEmail"))
	return ud.userRepository.FindUserByEmail(email)

}
