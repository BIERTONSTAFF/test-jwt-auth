package utils

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(l int) (string, error) {
	bytes := make([]byte, l)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	token := base64.StdEncoding.EncodeToString(bytes)

	return token, nil
}

func EncodeToken(token string) (string, error) {
	encodedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(encodedToken), nil
}

func CompareToken(hash string, token string) bool {
	eq := bcrypt.CompareHashAndPassword([]byte(hash), []byte(token))

	return eq == nil
}
