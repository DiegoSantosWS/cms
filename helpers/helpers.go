package helpers

import (
	"fmt"

	"github.com/DiegoSantosWS/cms/cone"
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

func GetNameGrupo(id int) (string, error) {
	var name string
	sql := "SELECT name FROM group WHERE id = ?"
	rows, err := cone.Db.Queryx(sql, id)
	if err != nil {
		fmt.Println("Erro: nome não encontrado", err.Error())
		return "", err
	}
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("Erro: nome não encontrado", err.Error())
			return "", err
		}
	}
	return string(name), nil
}
