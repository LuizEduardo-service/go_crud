package view

import (
	"github.com/LuizEduardo-service/go_crud/src/controller/model/response"
	"github.com/LuizEduardo-service/go_crud/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserReponse {
	return response.UserReponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
