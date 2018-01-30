package models

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/sessions"

	"github.com/DiegoSantosWS/cms/cone"
	"github.com/DiegoSantosWS/cms/helpers"

	ctrl "github.com/DiegoSantosWS/cms/controller"
)

var (
	key   = []byte("1234567890dhskai#sobn")
	Store = sessions.NewCookieStore(key)
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
	Senha   string `db:"pass"`
}

//Logout faz logout com usuario
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "logado")
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
	} else if r.Method == http.MethodPost {

		if r.FormValue("user") == "" && r.FormValue("pass") == "" {
			data := map[string]interface{}{
				"Msg": "Entre com dados de usuarios",
			}
			if err := ctrl.ModelosLogin.ExecuteTemplate(w, "login.html", data); err != nil {
				http.Error(w, "[TEMPLATE LOGIN], erro no carregamento", http.StatusInternalServerError)
				fmt.Println("LOGIN: erro na execução do modelo", err.Error())
			}
		}
		//Vamos receber os dados do formulario
		usr := UsuarioLogin{
			User: r.FormValue("user"),
			Pass: r.FormValue("pass"),
		}

		//Vamos verificar se usuario se senha Existe
		user := UsuarioLogado{}
		sql := "SELECT id, name, email, login, pass, type FROM users WHERE login =? "
		rows, err := cone.Db.Queryx(sql, usr.User)
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
		entrar := helpers.CheckPasswordHash(usr.Pass, user.Senha)
		if entrar != false {
			session, _ := Store.Get(r, "logado")
			session.Values["ID"] = user.ID
			session.Values["Nome"] = user.Nome
			session.Values["Email"] = user.Email
			session.Values["Tipo"] = user.Tipo
			session.Values["autorizado"] = true
			session.Save(r, w)
			CheckSession(w, r)
			//LogAcesso(user.Nome, user.Tipo, "Sucesso")
			http.Redirect(w, r, "/internal", 302)
		}
		//LogAcesso(user.Nome, user.Tipo, "Falha")
		http.Redirect(w, r, "/", 302)
	}
}

//CheckSession verifica uma sessão
func CheckSession(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "logado")
	if auth, ok := session.Values["autorizado"].(bool); !ok || !auth {
		http.Error(w, "Acesso negado", http.StatusInternalServerError)
		return
	}
}

//LogAcesso salva um log de acesso false ou true
func LogAcesso(usr, tipo, status string) {
	sql := "INSERT INTO log_access (user, tipo, status, data) VALUES (?,?,?, NOW() ) "
	_, err := cone.Db.Exec(sql, usr, tipo, status, time.Now().Format("02/01/2006 15:04:05"))
	if err != nil {
		fmt.Println("[LOG] error no log", err.Error())
	}
}
