package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

//HashPassword encripta uma senha passada para o bando de dados
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash compara uma senha e retona verdadeiro ou false
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
