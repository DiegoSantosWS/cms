package models

import (
	"fmt"
	"net/http"

	"github.com/DiegoSantosWS/cms/controller"
	"github.com/gorilla/mux"

	"github.com/DiegoSantosWS/cms/cone"
)

//DadosGrupos recebe os dados para listagem
type DadosGrupos struct {
	Id   int
	Nome string
}

//AltGrupos recebe os dados para alterar
type AltGrupos struct {
	ID   int
	Nome string
}

//Grupos Executa consulta e retornar uma slice para uma template
func Grupos(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r)
	//Busca todas informaçoes de grupos
	sql := "SELECT id, name FROM groups "
	rows, err := cone.Db.Queryx(sql)
	if err != nil {
		fmt.Println("[GRUPO] Erro, erro ao buscar lista com grupos")
		return
	}
	defer rows.Close()

	var g DadosGrupos
	var grs []DadosGrupos

	for rows.Next() {
		err := rows.Scan(&g.Id, &g.Nome)
		if err != nil {
			fmt.Println("[GRUPO] Erro, erro ao buscar lista com grupos")
			return
		}

		grs = append(grs, g)
	}

	data := map[string]interface{}{
		"Title":  "Lista ge grupos",
		"grupos": grs,
	}
	//Executando template
	if err := controller.ModelosGrupos.ExecuteTemplate(w, "listGroup.html", data); err != nil {
		http.Error(w, "[ERRO AO EXECUTAR A TEMPLATE]", http.StatusInternalServerError)
		fmt.Println("Error Grupos", err.Error())
	}
}

//NewGrupos Cadastra um novo grupo
func NewGrupos(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r) //Verifica se a sessão está iniciada
	if r.Method != http.MethodPost {
		controller.ModelosGruposN.Execute(w, nil)
		return
	}

	dados := AltGrupos{
		Nome: r.FormValue("name"),
	}

	sql := "INSERT INTO groups (name) VALUES (?) "
	stmt, err := cone.Db.Exec(sql, dados.Nome)
	if err != nil {
		fmt.Println("[CADGRUPO:] Erro na inclusão do grupo", sql, " - ", err.Error())
	}

	_, errs := stmt.RowsAffected()
	if errs != nil {
		fmt.Println("[CADGRUPO:] Erro ao pegar numero de linhas", sql, " - ", err.Error())
	}
	http.Redirect(w, r, "/grupos", 302)
	//fmt.Println("Linhas: ", linas, " linas(s) afetada(s)")
	//controller.ModelosGruposN.Execute(w, struct{ Secesso bool }{true})
}

//UpdateGrupos altera conteudo de um grupo
func UpdateGrupos(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r) //Verifica se a sessão está iniciada
	vars := mux.Vars(r)
	id := vars["id"]

	if r.Method != http.MethodPost {
		sql := "SELECT * FROM groups WHERE id = ? "
		linha, err := cone.Db.Queryx(sql, id)
		if err != nil {
			http.Error(w, "[ERROR:] Não foi encontrado nenhum grupos", http.StatusInternalServerError)
			fmt.Println("[ERROR] Nenhum grupo encontrado", err.Error())
			return
		}
		defer linha.Close()
		g := DadosGrupos{}
		for linha.Next() {
			err := linha.Scan(&g.Id, &g.Nome)
			if err != nil {
				http.Error(w, "[ERROR:] Não foi encontrado nenhum grupos", http.StatusInternalServerError)
				fmt.Println("[ERROR] Nenhum grupo encontrado", err.Error())
				return
			}
		}
		controller.ModelosGruposC.Execute(w, g)
		return
	}

	dados := AltGrupos{
		Nome: r.FormValue("name"),
	}

	sql := " UPDATE groups SET name = ? WHERE id = ? "
	line, err := cone.Db.Exec(sql, dados.Nome, id)
	if err != nil {
		http.Error(w, "[ERROR:]Não foi possível alterar grupo ", http.StatusInternalServerError)
		fmt.Println("[ERROR]Não foi possível alterar grupo ", err.Error())
	}
	_, err = line.RowsAffected()
	if err != nil {
		http.Error(w, "[ERROR:]Não foi possível alterar grupo ", http.StatusInternalServerError)
		fmt.Println("[ERROR]Não foi possível alterar grupo ", err.Error())
	}
	http.Redirect(w, r, "/grupos/"+id, 302)
}

//DeleteGrupos deleta um grupo
func DeleteGrupos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	methods := vars["method"]

	if methods == "delete" {
		sql, err := cone.Db.Queryx("DELETE FROM groups WHERE id = ?", id)
		if err != nil {
			http.Error(w, "[DELETE] erro ao delatar um grupo,", http.StatusInternalServerError)
			fmt.Println("[DELETE] erro ao deletar um grupo", sql, " - ", err.Error())
		}
		http.Redirect(w, r, "/grupos", 302)
	}
}
