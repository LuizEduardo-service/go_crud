package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/LuizEduardo-service/go_crud/src/configuration/rest_err"
	"github.com/golang-jwt/jwt"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {

	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"age":   ud.age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), //determina o tempo de expiração do token
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //definindo metodo de encripitação

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", rest_err.NewInternalServerError(fmt.Sprintf("erro ao tentar gerar a chave de acesso: %s", err.Error()))
	}
	return tokenString, nil
}

func VerifyToken(tokenValue string) (UserDomainInterface, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)
	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, rest_err.NewBadRequestError("token Invalido")
	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedError("token Invalido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedError("token Invalido")
	}

	return &userDomain{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   int8(claims["age"].(float64)),
	}, nil
}
func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix("Bearer ", token)
	}
	return token
}
