package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/DiegoSantosWS/cms/cone"
)

func ListContent(w http.ResponseWriter, r *http.Request) {
	sql := "SELECT c.id, c.title, c.description, c.date_ini, c.date_end, c.group, c.category_content  FROM content as c "
	rows, err := cone.Db.Queryx(sql)
	if err != nil {
		fmt.Println("[CONTEUDO] Erro ao buscar informações de conteudo: ", sql, " - ", err.Error())
		return
	}

	type Contents struct {
		Id        int    `json:"id"`
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
		var id int
		var title string
		var group string
		var description, date_ini, date_end, category_content string

		rows.Scan(&id, &title, &description, &date_ini, &date_end, &group, &category_content)
		g := GetNameGrupo(group)
		c := GetNameCategoria(category_content)
		contents = append(contents, Contents{id, title, description, date_ini, date_end, g, c})
	}
	contentData, err := json.Marshal(&contents)
	if err != nil {
		fmt.Println("[CONTEUDO] Erro ao buscar informações de conteudo: ", err.Error())
		return
	}
	w.Write(contentData)
}

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

func ListGroup(w http.ResponseWriter, r *http.Request) {
	sql := "SELECT g.id, g.name  FROM groups as g "
	rows, err := cone.Db.Queryx(sql)
	if err != nil {
		fmt.Println("[GRUPO] Erro ao buscar informações de GRUPO: ", sql, " - ", err.Error())
		return
	}

	type Groups struct {
		Id   int    `json:"id"`
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
		Id   int    `json:"id"`
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
		Id        int    `json:"id"`
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

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}

	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("Error writing buffer", err.Error())
		return err
	}

	//open file handle
	fmt.Println("FILE NAME: ", filename)
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error os in open file:", err.Error())
		return err
	}
	defer fh.Close()

	//copy file to directory
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		fmt.Println("Error os in copy file from directory:", err.Error())
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		fmt.Println("Error teste:", err.Error())
		return err
	}
	defer resp.Body.Close()

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error RESP BODY:", err.Error())
		return err
	}

	fmt.Println("Status: \n", resp.Status)
	fmt.Println("Body: \n", string(resp_body))
	return nil
}

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

	f, err := os.OpenFile("uploadfile/"+namenewFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	io.Copy(f, files)
	idcontent, err := strconv.Atoi(r.URL.Path[12:])
	if err != nil {
		fmt.Println(err.Error())
	}
	s := fmt.Sprintf("/conteudo/%d", idcontent)
	http.Redirect(w, r, s, 302)
}
