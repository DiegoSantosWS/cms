package models

import (
	"fmt"
	"net/http"

	"github.com/DiegoSantosWS/cms/helpers"

	"github.com/DiegoSantosWS/cms/cone"
	"github.com/DiegoSantosWS/cms/controller"
)

//CadUserEx cadastra um usuario externo
type CadUserEx struct {
	Nome  string
	Email string
	Login string
	Senha string
	Tipo  string
	Hash  string
}

//CadUserExternal Cadastra um usuario externo
func CadUserExternal(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		controller.ModelosCadastrar.Execute(w, nil)
		return
	}
	dados := CadUserEx{
		Nome:  r.FormValue("name"),
		Email: r.FormValue("email"),
		Login: r.FormValue("usuario"),
		Senha: r.FormValue("pass"),
		Tipo:  r.FormValue("tipo"),
		Hash:  "dlklsdkldksdkldklsksd",
	}
	dados.Senha, _ = helpers.HashPassword(dados.Senha)
	//tx := cone.Db.MustBegin()
	//tx.MustExec(tx.Rebind("INSERT INTO users (nome, email, login, pass, type, hash) VALUES (?, ?, ?, ?, ?, ?) "), dados.Nome, dados.Email, dados.Login, dados.Senha, dados.Tipo, dados.Hash)
	//err := tx.Commit()

	sql := "INSERT INTO users (name, email, login, pass, type, hash) VALUES (?, ?, ?, ?, ?, ?) "
	stmt, err := cone.Db.Exec(sql, dados.Nome, dados.Email, dados.Login, dados.Senha, dados.Tipo, dados.Hash)
	if err != nil {
		fmt.Println("[CADEX:] Erro na inclusão do usuario", sql, " - ", err.Error())
	}

	linas, errs := stmt.RowsAffected()
	if errs != nil {
		fmt.Println("[CADEX:] Erro ao pegar numero de linhas", sql, " - ", err.Error())
	}
	fmt.Println("Linhas: ", linas, " linas(s) afetada(s)")
	controller.ModelosCadastrar.Execute(w, struct{ Secesso bool }{true})
}
