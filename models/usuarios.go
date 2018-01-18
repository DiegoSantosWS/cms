package models

import (
	"fmt"
	"net/http"

	"github.com/DiegoSantosWS/cms/cone"
	ctrl "github.com/DiegoSantosWS/cms/controller"
	"github.com/gorilla/mux"
)

//DadosUsuarios armazena os dados do login recebido do banco...
type DadosUsuarios struct {
	Id      int
	Nome    string
	Email   string
	Usuario string
	Tipo    string
}

//Usuarios carrega template da lista dos usuarios
func Usuarios(w http.ResponseWriter, r *http.Request) {
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

//Usuario abre cadastro do usuario para realizar o update
func Usuario(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(r.URL.Path)

	if r.Method == http.MethodGet {
		fmt.Println(id)
	}
}

//DeleteUsuario deleta um usuario
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	meths := vars["method"]
	fmt.Println(meths)

	if meths == "delete" {
		fmt.Println("DELETE: ", id)
	}
}
