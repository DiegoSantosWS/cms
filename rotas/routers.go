package rotas

import (
	"net/http"

	"github.com/DiegoSantosWS/cms/models"
	"github.com/gorilla/mux"
)

func Routers() {
	r := mux.NewRouter()
	//Pegando rotas de arquivos estaticos
	fs := http.FileServer(http.Dir("assets/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	//Rotas de menus do sistema
	r.HandleFunc("/", models.Login)            //abre arquivo de login.html
	r.HandleFunc("/login", models.Login)       //realiza login do usuario
	r.HandleFunc("/internal", models.Internal) //realiza login do usuario
	r.HandleFunc("/usuarios", models.Usuarios)
	r.HandleFunc("/usuarios/{id}", models.Usuario)
	r.HandleFunc("/usuarios/{method}/{id}", models.DeleteUsuario)
	r.HandleFunc("/cad-user", models.CadUserExternal) //cadastra um usuario para entrar no sistema
	http.ListenAndServe(":8181", r)                   //inicia o servidor recebendo as rotas atravez do objeto r
}
