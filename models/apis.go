package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/DiegoSantosWS/cms/cone"
)

//ListContent lista conteudo cadastrados
func ListContent(w http.ResponseWriter, r *http.Request) {
	sql := "SELECT c.id, c.title, c.description, c.date_ini, c.date_end, c.group, c.category_content  FROM content as c "
	rows, err := cone.Db.Queryx(sql)
	if err != nil {
		fmt.Println("[CONTEUDO] Erro ao buscar informações de conteudo: ", sql, " - ", err.Error())
		return
	}

	type Contents struct {
		ID        int    `json:"id"`
		Titulo    string `json:"title"`
		Descricao string `json:"description"`
		DataIni   string `json:"date_ini"`
		DataFim   string `json:"date_end"`
		Grupo     string `json:"group"`
		Categoria string `json:"category_content"`
	}
	defer rows.Close()

	var contents []Contents
	for rows.Next() {
		var (
			id              int
			title           string
			group           string
			description     string
			dateIni         string
			dateEnd         string
			categoryContent string
		)

		rows.Scan(&id, &title, &description, &dateIni, &dateEnd, &group, &categoryContent)
		g := GetNameGrupo(group)
		c := GetNameCategoria(categoryContent)
		contents = append(contents, Contents{id, title, description, dateIni, dateEnd, g, c})
	}
	contentData, err := json.Marshal(&contents)
	if err != nil {
		fmt.Println("[CONTEUDO] Erro ao buscar informações de conteudo: ", err.Error())
		return
	}
	w.Write(contentData)
}

//GetNameGrupo retorna nome do grupo
func GetNameGrupo(id string) string {
	var name string
	sql := "SELECT name FROM groups WHERE id = ?"
	rows, err := cone.Db.Queryx(sql, id)
	if err != nil {
		fmt.Println("Erro: nome não encontrado", err.Error())
		return ""
	}
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("Erro: nome não encontrado", err.Error())
			return ""
		}
	}
	return string(name)
}

//GetNameCategoria retorna nome da categoria
func GetNameCategoria(id string) string {
	var name string
	sql := "SELECT categoria as name FROM categorys WHERE id = ?"
	rows, err := cone.Db.Queryx(sql, id)
	if err != nil {
		fmt.Println("Erro: nome não encontrado", err.Error())
		return ""
	}
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("Erro: nome não encontrado", err.Error())
			return ""
		}
	}
	return string(name)
}

//ListGroup retorna lista das categorias
func ListGroup(w http.ResponseWriter, r *http.Request) {
	sql := "SELECT g.id, g.name  FROM groups as g "
	rows, err := cone.Db.Queryx(sql)
	if err != nil {
		fmt.Println("[GRUPO] Erro ao buscar informações de GRUPO: ", sql, " - ", err.Error())
		return
	}

	type Groups struct {
		ID   int    `json:"id"`
		Nome string `json:"name"`
	}
	defer rows.Close()

	var groups []Groups
	for rows.Next() {
		var id int
		var name string

		rows.Scan(&id, &name)

		groups = append(groups, Groups{id, name})
	}
	groupData, err := json.Marshal(&groups)
	if err != nil {
		fmt.Println("[GRUPO] Erro ao buscar informações de GRUPO: ", err.Error())
		return
	}
	w.Write(groupData)
}

//ListCategorysByGroup LISTA CATEGORIAS POR GRUPO
func ListCategorysByGroup(w http.ResponseWriter, r *http.Request) {
	idgroup, err := strconv.Atoi(r.URL.Path[26:])
	if err != nil {
		http.Error(w, "Não foi enviado um codigo valido.", http.StatusBadRequest)
		fmt.Println("[GRUPO] Erro id não econtrado", err.Error())
		return
	}

	sql := "SELECT c.id, c.categoria  FROM categorys as c WHERE grupo = ? "
	rows, err := cone.Db.Queryx(sql, idgroup)
	if err != nil {
		fmt.Println("[CATEGORIA] Erro ao buscar informações de GRUPO: ", sql, " - ", err.Error())
		return
	}

	type Categorys struct {
		ID   int    `json:"id"`
		Nome string `json:"categoria"`
	}
	defer rows.Close()

	var cats []Categorys
	for rows.Next() {
		var id int
		var categoria string

		rows.Scan(&id, &categoria)

		cats = append(cats, Categorys{id, categoria})
	}
	categorysData, err := json.Marshal(&cats)
	if err != nil {
		fmt.Println("[GRUPO] Erro ao buscar informações de GRUPO: ", err.Error())
		return
	}
	w.Write(categorysData)
}

//ListContentByID LISTA CONTEUDO PELO ID
func ListContentByID(w http.ResponseWriter, r *http.Request) {

	idcontent, err := strconv.Atoi(r.URL.Path[21:])
	if err != nil {
		http.Error(w, "Não foi enviado um codigo valido.", http.StatusBadRequest)
		fmt.Println("[GRUPO] Erro id não econtrado", err.Error())
		return
	}

	sql := "SELECT c.id, c.title, c.description, c.date_ini, c.date_end, c.group, c.category_content, c.text  FROM content as c WHERE id = ?"
	rows, err := cone.Db.Queryx(sql, idcontent)
	if err != nil {
		fmt.Println("[CONTEUDO] Erro ao buscar informações de conteudo: ", sql, " - ", err.Error())
		return
	}

	type Conts struct {
		ID        int    `json:"id"`
		Titulo    string `json:"title"`
		Descricao string `json:"description"`
		DataIni   string `json:"date_ini"`
		DataFim   string `json:"date_end"`
		Grupo     string `json:"group"`
		Categoria string `json:"category_content"`
		Texto     string `json:"text"`
	}
	defer rows.Close()

	var cont []Conts
	for rows.Next() {
		var id int
		var title string
		var group string
		var description, date_ini, date_end, category_content, text string

		rows.Scan(&id, &title, &description, &date_ini, &date_end, &group, &category_content, &text)
		cont = append(cont, Conts{id, title, description, date_ini, date_end, group, category_content, text})
	}
	contData, err := json.Marshal(&cont)
	if err != nil {
		fmt.Println("[CONTEUDO] Erro ao buscar informações de conteudo: ", err.Error())
		return
	}
	w.Write(contData)
}

//DeleteContent deleta conteudo conforme id informado
func DeleteContent(w http.ResponseWriter, r *http.Request) {

	var status int
	var menssage string

	idcontent, err := strconv.Atoi(r.URL.Path[19:])
	if err != nil {
		http.Error(w, "Codigo não informado", http.StatusInternalServerError)
		fmt.Println("Codigo não encontrado", err.Error())
		status = 404
		menssage = "Codigo não encontrado"
	}

	type Msgs struct {
		Status   int    `json:"status"`
		Menssage string `json:"menssage"`
	}
	var cats []Msgs

	if idcontent > 0 {
		sql, err := cone.Db.Queryx("DELETE FROM content WHERE id = ?", idcontent)
		if err != nil {
			http.Error(w, "[DELETE] erro ao delatar um conteudo ,", http.StatusInternalServerError)
			fmt.Println("[DELETE] erro ao deletar um conteudo", sql, " - ", err.Error())
			status = 301
			menssage = "Codigo encontrado encontrado"
		}
		status = 302
		menssage = "Codigo encontrado encontrado"
	}
	cats = append(cats, Msgs{status, menssage})
	delete, err := json.Marshal(&cats)
	if err != nil {
		fmt.Println("[GRUPO] Erro ao buscar informações de GRUPO: ", err.Error())
		return
	}
	w.Write(delete)
}

func ListFileContent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[21:])
	if err != nil {
		http.Error(w, "Não foi enviado um codigo valido.", http.StatusBadRequest)
		fmt.Println("[GRUPO] Erro id não econtrado", err.Error())
		return
	}

	sql := "SELECT c.id, c.nome, c.path  FROM ged as c WHERE idref = ?"
	rows, err := cone.Db.Queryx(sql, id)
	if err != nil {
		fmt.Println("[CONTEUDO] Erro ao buscar informações de arquivos: ", sql, " - ", err.Error())
		return
	}

	type Ged struct {
		ID      int    `json:"id"`
		Nome    string `json:"nome"`
		Caminho string `json:"path"`
	}
	defer rows.Close()

	var ged []Ged
	for rows.Next() {
		var id int
		var nome string
		var path string

		rows.Scan(&id, &nome, &path)
		ged = append(ged, Ged{id, nome, path})
	}
	gedData, err := json.Marshal(&ged)
	if err != nil {
		fmt.Println("[CONTEUDO] Erro ao buscar informações de conteudo: ", err.Error())
		return
	}
	w.Write(gedData)
}

//Upload faz uploade para um diretorios
func Upload(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(32 << 20)
	files, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer files.Close()
	//fmt.Fprintf(w, "%v", handler.Header)
	//Novo nome do arquivo
	namenewFile := time.Now().Format("02012006150405") + handler.Filename

	f, err := os.OpenFile("assets/uploadfile/"+namenewFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	io.Copy(f, files)
	id := r.FormValue("codigo")
	nome := namenewFile
	caminho := "uploadfile/" + namenewFile
	sql := "INSERT INTO ged (`nome`, `idref`, `path`) VALUES (?,?,?) "
	stmt, err := cone.Db.Exec(sql, nome, id, caminho)
	if err != nil {
		fmt.Println("[CADGED:] Erro na inclusão da imagem ", sql, " - ", err.Error())
	}
	_, errs := stmt.RowsAffected()
	if errs != nil {
		fmt.Println("[CADGED:] Erro ao pegar numero de linhas", sql, " - ", err.Error())
	}
	s := fmt.Sprintf("/conteudo/%s", id)
	http.Redirect(w, r, s, 302)
}
