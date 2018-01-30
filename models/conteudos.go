package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/DiegoSantosWS/cms/controller"

	"github.com/DiegoSantosWS/cms/cone"
)

type CadContent struct {
	titulo    string
	descricao string
	dataI     string
	dataF     string
	grupo     string
	categoria string
	texto     string
	slug      string
}
type AltContent struct {
	titulo    string
	descricao string
	dataI     string
	dataF     string
	grupo     string
	categoria string
	texto     string
	slug      string
}

//Conteudos carregando informações de conteudos
func Conteudos(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r)
	if r.Method == http.MethodPost {
		dados := CadContent{

			titulo:    r.FormValue("tituloContent"),
			descricao: r.FormValue("descContent"),
			dataI:     r.FormValue("dateIni"),
			dataF:     r.FormValue("dateEnd"),
			grupo:     r.FormValue("group"),
			categoria: r.FormValue("categoriaContent"),
			texto:     r.FormValue("texto"),
			slug:      "teste-teste",
		}

		sql := "INSERT INTO content (`title`, `description`, `text`, `slug`, `date_ini`, `date_end`, `group`, `category_content`) VALUES (?,?,?,?,?,?,?,?) "
		stmt, err := cone.Db.Exec(sql, dados.titulo, dados.descricao, dados.texto, dados.slug, dados.dataI, dados.dataF, dados.grupo, dados.categoria)
		if err != nil {
			fmt.Println("[CADCONTENT:] Erro na inclusão do conteudo", sql, " - ", err.Error())
		}

		_, errs := stmt.RowsAffected()
		if errs != nil {
			fmt.Println("[CADGRUPO:] Erro ao pegar numero de linhas", sql, " - ", err.Error())
		}
		cod, _ := stmt.LastInsertId()
		s := fmt.Sprintf("/conteudo/%d", cod)
		http.Redirect(w, r, s, 302)
	}

	data := map[string]interface{}{
		"Title": "Conteúdos",
	}
	if err := controller.ModelosConteudo.ExecuteTemplate(w, "listConteudo.html", data); err != nil {
		http.Error(w, "[CONTENT ERRO] Erro in the execute template", http.StatusInternalServerError)
		fmt.Println("Erro ao executar template", err.Error())
	}
}

//UpdateConteudos alterando informações de conteudos
func UpdateConteudos(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r)
	id, err := strconv.Atoi(r.URL.Path[10:])
	if err != nil {
		http.Error(w, "Não foi enviado um codigo valido.", http.StatusBadRequest)
		fmt.Println("[GRUPO] Erro id não econtrado", err.Error())
		return
	}
	if r.Method == http.MethodPost {
		dados := AltContent{
			titulo:    r.FormValue("tituloContent"),
			descricao: r.FormValue("descContent"),
			dataI:     r.FormValue("dateIni"),
			dataF:     r.FormValue("dateEnd"),
			grupo:     r.FormValue("group"),
			categoria: r.FormValue("categoriaContent"),
			texto:     r.FormValue("texto"),
			slug:      "teste-teste",
		}

		sql := "UPDATE content SET `title`=?, `description`=?, `text`=?, `slug`=?, `date_ini`=?, `date_end`=?, `group`=?, `category_content`=? WHERE id = ?"
		stmt, err := cone.Db.Exec(sql, dados.titulo, dados.descricao, dados.texto, dados.slug, dados.dataI, dados.dataF, dados.grupo, dados.categoria, id)
		if err != nil {
			fmt.Println("[CADCONTENT:] Erro na inclusão do conteudo", sql, " - ", err.Error())
		}
		_, errs := stmt.RowsAffected()
		if errs != nil {
			http.Error(w, "[ERROR:]Não foi possível alterar grupo ", http.StatusInternalServerError)
			fmt.Println("[ERROR]Não foi possível alterar grupo ", errs.Error())
		}
		s := fmt.Sprintf("/conteudo/%d", id)
		http.Redirect(w, r, s, 302)

	}
	data := map[string]interface{}{
		"Title": "Conteúdos",
	}
	if err := controller.ModelosConteudoC.ExecuteTemplate(w, "altConteudo.html", data); err != nil {
		http.Error(w, "[CONTENT ERRO] Erro in the execute template", http.StatusInternalServerError)
		fmt.Println("Erro ao executar template", err.Error())
	}
}

//NewConteudos cadastrando informações de conteudos
func NewConteudos(w http.ResponseWriter, r *http.Request) {
	CheckSession(w, r)
	data := map[string]interface{}{
		"Title": "Conteúdos",
	}
	if err := controller.ModelosConteudoN.ExecuteTemplate(w, "newConteudo.html", data); err != nil {
		http.Error(w, "[CONTENT ERRO] Erro in the execute template", http.StatusInternalServerError)
		fmt.Println("Erro ao executar template", err.Error())
	}
}

//DeleteConteudos excluindo informações de conteudos
func DeleteConteudos(w http.ResponseWriter, r *http.Request) {

}

func listConteudo(sqlString string) (string, error) {

	rows, err := cone.Db.Queryx(sqlString)
	if err != nil {
		fmt.Println("[CONTEUDO] Erro ao buscar informações de conteudo: ", sqlString, " - ", err.Error())
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println("[CONTEUDO] Erro ao buscar informações da coluna: ", err.Error())
		return "", err
	}
	count := len(columns)
	tableDate := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuesPtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuesPtrs[i] = &values[i]
		}
		rows.Scan(valuesPtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableDate = append(tableDate, entry)
	}
	jsonData, err := json.Marshal(tableDate)
	if err != nil {
		fmt.Println("[CONTEUDO] Erro ao buscar informações de conteudo: ", sqlString, " - ", err.Error())
		return "", err
	}
	return string(jsonData), nil
}
