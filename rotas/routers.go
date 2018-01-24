package rotas

import (
	"net/http"

	"github.com/DiegoSantosWS/cms/models"
	"github.com/gorilla/mux"
)

//Routers - função para iniciar uma rota no sistema
func Routers() {
	r := mux.NewRouter()
	//Pegando rotas de arquivos estaticos
	fs := http.FileServer(http.Dir("assets/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	//Rotas de menus do sistema
	r.HandleFunc("/", models.Login)            //abre arquivo de login.html
	r.HandleFunc("/login", models.Login)       //realiza login do usuario
	r.HandleFunc("/logout", models.Logout)     //realiza logout do usuario
	r.HandleFunc("/internal", models.Internal) //leva para pagina princiapl do sistema

	r.HandleFunc("/usuarios", models.Usuarios)                    //abre a lista de usuarios
	r.HandleFunc("/usuarios/{id}", models.UpdateUsuario)          //update de usuario
	r.HandleFunc("/usuarios/{method}/{id}", models.DeleteUsuario) //deleta um usuario usuarios
	r.HandleFunc("/cad-user", models.CadUserExternal)             //cadastra um usuario para entrar no sistema

	r.HandleFunc("/grupos", models.Grupos)                     //Abre lista de grupos
	r.HandleFunc("/grupos/{id}", models.UpdateGrupos)          //Abre cadastro de grupo para editar
	r.HandleFunc("/newgrupo", models.NewGrupos)                //Abre formulario para cadastrar um novo grupo
	r.HandleFunc("/grupos/{method}/{id}", models.DeleteGrupos) //Deleta um grupo

	r.HandleFunc("/categorias", models.Categorias)                     //Abre lista de categorias
	r.HandleFunc("/categorias/{id}", models.UpdateCategorias)          //Abre cadastro de categoria para editar
	r.HandleFunc("/newcategorias", models.NewCategorias)               //Abre formulario para cadastrar um nova categoria
	r.HandleFunc("/categorias/{method}/{id}", models.DeleteCategorias) //Deleta uma categoria

	http.ListenAndServe(":12345", r) //inicia o servidor recebendo as rotas atravez do objeto r
}
