package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["expiresIn"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.SecretKey))
}

func ValidateToken(r *http.Request) error {
	tokenString := getTokenFromHeaders(r)

	token, erro := jwt.Parse(tokenString, validateSigningMethod)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token must be a valid token")
}

func getUserIdFromToken(r *http.Request) (uint64, error) {
	tokenString := getTokenFromHeaders(r)

	token, erro := jwt.Parse(tokenString, validateSigningMethod)
	if erro != nil {
		return 0, erro
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		/*
			permissions["userId"] pelo padrão do jwt é salvo no formato float.
			Aqui, é necessário convertê-lo para string antes de passá-lo como parâmetro
			para o ParseUint, que irá retornar o userId como uint64
		*/
		userId, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return userId, nil
	}

	return 0, errors.New("token must be a valid token")
}

func getTokenFromHeaders(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 { // verifica se o token possui a estrutura 'Bearer token'
		return strings.Split(token, " ")[1]
	}
	return token
}

func validateSigningMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
