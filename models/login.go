package models

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/DiegoSantosWS/cms/cone"

	ctrl "github.com/DiegoSantosWS/cms/controller"
)

var (
	key   = []byte("1234567890dhskai#sobn")
	store = sessions.NewCookieStore(key)
)

//UsuarioLogin struct para armazenar os dados do login
type UsuarioLogin struct {
	User string
	Pass string
}

//UsuarioLogado armazena os dados do login recebido do banco...
type UsuarioLogado struct {
	ID      int    `db:"id"`
	Nome    string `db:"name"`
	Email   string `db:"email"`
	Usuario string `db:"login"`
	Tipo    string `db:"type"`
}

//Logout faz logout com usuario
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "logado")
	session.Values["autorizado"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}

//Login carrega uma template
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		if err := ctrl.ModelosLogin.ExecuteTemplate(w, "login.html", nil); err != nil {
			http.Error(w, "[TEMPLATE LOGIN], erro no carregamento", http.StatusInternalServerError)
			fmt.Println("LOGIN: erro na execução do modelo", err.Error())
		}
		return
	}
	//Vamos receber os dados do formulario
	usr := UsuarioLogin{
		User: r.FormValue("user"),
		Pass: r.FormValue("pass"),
	}

	//Vamos verificar se usuario se senha Existe
	user := UsuarioLogado{}
	sql := "SELECT id, name, email, login FROM users WHERE login =? AND pass = ?"
	rows, err := cone.Db.Queryx(sql, usr.User, usr.Pass)
	if err != nil {
		fmt.Println("[LOGIN] Erro ao executar o login", sql, " - ", err.Error())
		return
	}
	defer rows.Close()
	//for para buscar os dados do usuario
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			fmt.Println("[LOGIN] Erro ao executar o login", sql, " - ", err.Error())
			return
		}
	}
	session, _ := store.Get(r, "logado")
	session.Values["ID"] = user.ID
	session.Values["Nome"] = user.Nome
	session.Values["Email"] = user.Email
	session.Values["Tipo"] = user.Tipo
	session.Values["autorizado"] = true
	session.Save(r, w)
	http.Redirect(w, r, "/internal", 302)
}
