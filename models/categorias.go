package models

import (
	"fmt"
	"net/http"

	"github.com/DiegoSantosWS/cms/controller"
	"github.com/gorilla/mux"

	"github.com/DiegoSantosWS/cms/cone"
)

//DadosCategorias recebe os dados para listagem
type DadosCategorias struct {
	Id        int
	Grupo     string
	Nome      string
	NomeGrupo string
}

//AltCategorias recebe os dados para alterar
type AltCategorias struct {
	ID        int
	Group     string
	Nome      string
	NomeGrupo string
}

//Categorias Executa consulta e retornar uma slice para uma template
func Categorias(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r)
	//Busca todas informaçoes de grupos
	sql := "SELECT c.id, c.grupo, c.categoria, g.name FROM categorys as c INNER JOIN groups as g ON g.id = c.grupo "
	rows, err := cone.Db.Queryx(sql)
	if err != nil {
		fmt.Println("[Categorias] Erro, erro ao buscar lista com categorias")
		return
	}
	defer rows.Close()

	var c DadosCategorias
	var crs []DadosCategorias

	for rows.Next() {
		err := rows.Scan(&c.Id, &c.Grupo, &c.Nome, &c.NomeGrupo)
		if err != nil {
			fmt.Println("[Categorias] Erro, erro ao buscar lista com grupos")
			return
		}

		crs = append(crs, c)
	}

	data := map[string]interface{}{
		"Title":      "Lista de Categorias",
		"categorias": crs,
	}
	//Executando template
	if err := controller.ModelosCategorias.ExecuteTemplate(w, "listCategoria.html", data); err != nil {
		http.Error(w, "[ERRO AO EXECUTAR A TEMPLATE]", http.StatusInternalServerError)
		fmt.Println("Error Grupos", err.Error())
	}
}

//NewCategorias Cadastra um novo grupo
func NewCategorias(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r) //Verifica se a sessão está iniciada
	if r.Method != http.MethodPost {
		sql := "SELECT id, name FROM groups "
		rows, err := cone.Db.Queryx(sql)
		if err != nil {
			fmt.Println("[Categorias] Erro, erro ao buscar lista com categorias")
			return
		}
		defer rows.Close()

		var c grupos
		var crs []grupos

		for rows.Next() {
			err := rows.Scan(&c.Id, &c.Nome)
			if err != nil {
				fmt.Println("[Categorias] Erro, erro ao buscar lista com grupos", err.Error())
				return
			}

			crs = append(crs, c)
		}
		data := map[string]interface{}{
			"grupos": crs,
		}
		controller.ModelosCategoriasN.Execute(w, data)
		return
	}

	dados := AltCategorias{
		Nome:  r.FormValue("categoria"),
		Group: r.FormValue("group"),
	}

	sql := "INSERT INTO categorys (grupo, categoria) VALUES (?, ?) "
	stmt, err := cone.Db.Exec(sql, dados.Group, dados.Nome)
	if err != nil {
		fmt.Println("[CADGRUPO:] Erro na inclusão da categoria", sql, " - ", err.Error())
	}

	_, errs := stmt.RowsAffected()
	if errs != nil {
		fmt.Println("[CADGRUPO:] Erro ao pegar numero de linhas", sql, " - ", err.Error())
	}
	http.Redirect(w, r, "/categorias", 302)
	//fmt.Println("Linhas: ", linas, " linas(s) afetada(s)")
	//controller.ModelosGruposN.Execute(w, struct{ Secesso bool }{true})
}

//UpdateCategorias altera conteudo de um grupo
func UpdateCategorias(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r) //Verifica se a sessão está iniciada
	vars := mux.Vars(r)
	id := vars["id"]

	if r.Method != http.MethodPost {
		sql := "SELECT c.id, c.grupo, c.categoria, g.name FROM categorys as c INNER JOIN groups as g ON c.grupo = g.id WHERE c.id = ? "
		linha, err := cone.Db.Queryx(sql, id)
		if err != nil {
			http.Error(w, "[ERROR:] Não foi encontrado nenhum categoria", http.StatusInternalServerError)
			fmt.Println("[ERROR] Nenhum categoria encontrado", sql, " - ", err.Error())
			return
		}
		defer linha.Close()
		g := DadosCategorias{}
		for linha.Next() {
			err := linha.Scan(&g.Id, &g.Grupo, &g.NomeGrupo, &g.Nome)
			if err != nil {
				http.Error(w, "[ERROR:] Não foi encontrado nenhuma categoria", http.StatusInternalServerError)
				fmt.Println("[ERROR] Nenhum categoria encontrado", err.Error())
				return
			}
		}

		sql = "SELECT id, name FROM groups "
		rows, err := cone.Db.Queryx(sql)
		if err != nil {
			fmt.Println("[Categorias] Erro, erro ao buscar lista com categorias")
			return
		}
		defer rows.Close()

		var c grupos
		var crs []grupos

		for rows.Next() {
			err := rows.Scan(&c.Id, &c.Nome)
			if err != nil {
				fmt.Println("[Categorias] Erro, erro ao buscar lista com grupos", err.Error())
				return
			}

			crs = append(crs, c)
		}
		data := map[string]interface{}{
			"categorias": g,
			"grupos":     crs,
		}

		controller.ModelosCategoriasC.Execute(w, data)
		return
	}

	dados := AltCategorias{
		Nome:  r.FormValue("categoria"),
		Group: r.FormValue("group"),
	}

	sql := " UPDATE categorys SET grupo = ?, categoria = ? WHERE id = ? "
	line, err := cone.Db.Exec(sql, dados.Group, dados.Nome, id)
	if err != nil {
		http.Error(w, "[ERROR:]Não foi possível alterar grupo ", http.StatusInternalServerError)
		fmt.Println("[ERROR]Não foi possível alterar grupo ", err.Error())
	}
	_, err = line.RowsAffected()
	if err != nil {
		http.Error(w, "[ERROR:]Não foi possível alterar grupo ", http.StatusInternalServerError)
		fmt.Println("[ERROR]Não foi possível alterar grupo ", err.Error())
	}
	http.Redirect(w, r, "/categorias/"+id, 302)
}

//DeleteCategorias deleta um grupo
func DeleteCategorias(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r)
	vars := mux.Vars(r)
	id := vars["id"]
	methods := vars["method"]

	if methods == "delete" {
		sql, err := cone.Db.Queryx("DELETE FROM categorys WHERE id = ?", id)
		if err != nil {
			http.Error(w, "[DELETE] erro ao delatar uma categoria ,", http.StatusInternalServerError)
			fmt.Println("[DELETE] erro ao deletar uma categoria", sql, " - ", err.Error())
		}
		http.Redirect(w, r, "/categorias", 302)
	}
}

type grupos struct {
	Id   int
	Nome string
}
