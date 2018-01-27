# cms
CMS USANDO Go!

Sistema é apenas um gerenciador de conteudo para sites com poucas paginas
Sistema ainda está em Desenvolvimento por tanto ainda não colocarei explicação de seu funcinamento completo


<h1><b>Como funciona</b></h1>
<ul>
  <li> 1 - Criar banco de dados e fazer as conexões devidas de usuário e senha</li>
  <li> 2 - Criar um usuário e senha do tipo Admin</li>
  <li> 3 - So começar a usar no seu site.</li>
</ul>
<h1><b>ENVIANDO IMAGEM</b></h1>
<code>
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

</code>