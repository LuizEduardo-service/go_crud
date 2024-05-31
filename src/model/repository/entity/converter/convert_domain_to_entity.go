package converter

import (
	"github.com/LuizEduardo-service/go_crud/src/model"
	"github.com/LuizEduardo-service/go_crud/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}

}
