package model

import "github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string
	GetID() string

	GenerateToken() (string, *rest_err.RestErr)
	SetID(string)
	EncryptPassword()
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func NewUserUpdateDomain(
	name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}

func NewUserLoginDomain(
	email,
	password string,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}
