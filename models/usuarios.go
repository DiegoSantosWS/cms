package models

import (
	"fmt"
	"net/http"

	"github.com/DiegoSantosWS/cms/controller"
	"github.com/DiegoSantosWS/cms/helpers"

	"github.com/DiegoSantosWS/cms/cone"
	ctrl "github.com/DiegoSantosWS/cms/controller"
	"github.com/gorilla/mux"
)

//DadosUsuarios armazena os dados do login recebido do banco...
type DadosUsuarios struct {
	Id        int
	Nome      string
	Email     string
	Usuario   string
	Tipo      string
	Selectded bool
}

//AltUserEx - altera dados do usuario
type AltUserEx struct {
	Nome  string
	Email string
	Login string
	Senha string
	Tipo  string
	Hash  string
}

//Usuarios carrega template da lista dos usuarios
func Usuarios(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r)
	sql := "SELECT id, name, email, login, type FROM users "
	rows, err := cone.Db.Queryx(sql)
	if err != nil {
		fmt.Println("[LOGIN] Erro ao executar o login", sql, " - ", err.Error())
		return
	}
	defer rows.Close()
	var u DadosUsuarios
	var usr []DadosUsuarios
	//for para buscar os dados do usuario
	for rows.Next() {
		err := rows.Scan(&u.Id, &u.Nome, &u.Email, &u.Usuario, &u.Tipo)
		if err != nil {
			fmt.Println("[LOGIN] Erro ao executar o login", sql, " - ", err.Error())
			return
		}

		usr = append(usr, u)
	}

	data := map[string]interface{}{
		"titulo":   "Listagem de usuarios",
		"usuarios": usr,
	}

	if err := ctrl.ModelosUsuarioslist.ExecuteTemplate(w, "listUser.html", data); err != nil {
		http.Error(w, "[TEMPLATE LOGIN], erro no carregamento", http.StatusInternalServerError)
		fmt.Println("LOGIN: erro na execução do modelo", err.Error())
	}
}

//UpdateUsuario abre cadastro do usuario para realizar o update
func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r) //Verifica se a sessão está iniciada
	vars := mux.Vars(r)
	id := vars["id"]
	//Veifica o metodo inicial para não alterar e sim bucar as informação
	if r.Method != http.MethodPost {

		sql := "SELECT id, name, email, login, type FROM users WHERE id = ? "
		linha, err := cone.Db.Queryx(sql, id)
		if err != nil {
			http.Error(w, "[ERRO] Código inixistente. verifique.", http.StatusInternalServerError)
			fmt.Println("[ERRO] Usuário não encontrado", err.Error())
			return
		}
		defer linha.Close()
		u := DadosUsuarios{}
		for linha.Next() {
			err := linha.Scan(&u.Id, &u.Nome, &u.Email, &u.Usuario, &u.Tipo)
			if err != nil {
				http.Error(w, "[ERRO] Usuário não encontrado", http.StatusInternalServerError)
				fmt.Println("[ERRO] Usuário não encontrado", err.Error())
				return
			}
			if u.Tipo != "Admin" {
				u.Selectded = true
			} else {
				u.Selectded = true
			}
		}
		controller.ModelosUsuariosPUT.Execute(w, u)
		return
	}
	//Passa para struct os metodos recebido do post
	dados := AltUserEx{
		Nome:  r.FormValue("name"),
		Email: r.FormValue("email"),
		Login: r.FormValue("usuario"),
		Senha: r.FormValue("pass"),
		Tipo:  r.FormValue("tipo"),
		Hash:  "dsasdçdaskdadk",
	}
	var cpSenha string
	if r.FormValue("pass") != "" {
		cpSenha, _ = helpers.HashPassword(r.FormValue("pass"))
	} else { //Fiz assim pq não sei outra forma usando a lingagem GO! se alguem souber pode me mandar por e-mail: tec.infor321@gmail.com
		sql := "SELECT pass FROM users WHERE id = ? "
		linha, err := cone.Db.Queryx(sql, id)
		if err != nil {
			http.Error(w, "[ERRO] Código inixistente. verifique.", http.StatusInternalServerError)
			fmt.Println("[ERRO] Usuário não encontrado", err.Error())
			return
		}
		defer linha.Close()
		var senha string
		for linha.Next() {
			err := linha.Scan(&senha)
			if err != nil {
				http.Error(w, "[ERRO] Usuário não encontrado", http.StatusInternalServerError)
				fmt.Println("[ERRO] Usuário não encontrado", err.Error())
				return
			}
			cpSenha = senha
		}
	}
	sql := "UPDATE users SET name = ?, email = ?, login = ?, pass = ?, type = ? WHERE id = ?;"
	line, err := cone.Db.Exec(sql, dados.Nome, dados.Email, dados.Login, cpSenha, dados.Tipo, id)
	if err != nil {
		http.Error(w, "[UPDATE USER] Erro na alteração, ", http.StatusInternalServerError)
		fmt.Println("[UPDATE USER] Não foi possivel alterar o usuario, ", err.Error())
	}

	_, err = line.RowsAffected()
	if err != nil {
		http.Error(w, "[UPDATE USER] Não alterou nenhuma linha, ", http.StatusInternalServerError)
		fmt.Println("[UPDATE USER] Não foi possivel alterar o usuario, ", err.Error())
	}
	http.Redirect(w, r, "/usuarios/"+id, 302)
}

//DeleteUsuario deleta um usuario
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r)
	vars := mux.Vars(r)
	id := vars["id"]
	meths := vars["method"]

	if meths == "delete" {
		//Executa comando
		sql, err := cone.Db.Queryx("DELETE FROM users WHERE id = ?", id)
		if err != nil {
			http.Error(w, "[DELETE], erro ao deletar", http.StatusInternalServerError)
			fmt.Println("DELETE: erro na execução do modelo", sql, " - ", err.Error())
		}
		http.Redirect(w, r, "/usuarios", 302)
	}
}
