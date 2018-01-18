package helpers

import (
	"net/http"

	m "github.com/DiegoSantosWS/cms/models"
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

func CheckSession(w http.ResponseWriter, r *http.Request) {
	session, _ := m.Store.Get(r, "logado")
	if auth, ok := session.Values["autorizado"].(bool); !ok || !auth {
		http.Error(w, "Acesso negado", http.StatusForbidden)
		return
	}
}
