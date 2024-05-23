package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type userDomain struct {
	ID       string
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
}

func (ud *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		return "", err
	}
	return string(b), nil

}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}

type UserDomainInterface interface {
	GetEmail() string
	GetPassaword() string
	GetAge() int8
	GetName() string

	SetID(string)
	GetJSONValue() (string, error)
	EncryptPassword()
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}
func (ud *userDomain) GetPassaword() string {
	return ud.Password
}
func (ud *userDomain) GetName() string {
	return ud.Name
}
func (ud *userDomain) GetAge() int8 {
	return ud.Age
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		"", email, password, name, age,
	}
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString((hash.Sum(nil)))
}
