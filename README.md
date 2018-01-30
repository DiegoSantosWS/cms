# cms
CMS USING Go!

System is just a content manager for sites with few pages.
The installation part is being developed.

<h1><b>How it works</b></h1>
<ul>
  <li> 1 - Create database and make the necessary user and password connections.</li>
  <li> 2 - Create a username and password of type Admin.</li>
  <li> 3 - <a href="#install">Installation</a>.</li>
  <li> 4 - Just start using it on your website.</li>
</ul>

<h1><b> <a href="#installation">Installation</a></b></h1>
<div id="install" style="color:red;">
	<p style="color:red;">By default system runs on the port http://localhost:3000/</p>
	<p>After creating the database, <a href="https://github.com/DiegoSantosWS/cms/blob/master/cone/conexao.go#L15">access the file conexao.go</a> and change connection data for your username and password</p>
	<div id='installation'>
		<p style="color:red;">URL: http://localhost:3000/install</p>
		<p>
		After performing the install, set whether the system created the tables correctly, if not,
		you have created copy the code that is in the root and run in your database.
		after that, just access, the<b> URL: http://localhost:3000/</b> enter your username and password.
		</p>
	</div>
</div>

<h1><b>Sending File</b></h1>
<p><a href="https://github.com/DiegoSantosWS/cms/blob/master/models/apis.go#L256">Example</a></p>
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
		//New name from file
		namenewFile := time.Now().Format("02012006150405") + handler.Filename

		f, err := os.OpenFile("assets/uploadfile/"+namenewFile, os.O_WRONLY|os.O_CREATE, 0666)
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