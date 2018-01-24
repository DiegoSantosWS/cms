package controller

import (
	"html/template"
)

var (
	ModelosLogin        = template.Must(template.ParseFiles("view/login.html"))
	ModelosCadastrar    = template.Must(template.ParseFiles("view/cad-user.html"))
	ModelosInternal     = template.Must(template.ParseFiles("view/interno.html"))
	ModelosUsuarioslist = template.Must(template.ParseFiles("view/usuarios/listUser.html"))
	ModelosUsuariosPUT  = template.Must(template.ParseFiles("view/usuarios/AltUser.html"))

	ModelosGrupos  = template.Must(template.ParseFiles("view/grupos/listGroup.html"))
	ModelosGruposC = template.Must(template.ParseFiles("view/grupos/newGroup.html"))
	ModelosGruposN = template.Must(template.ParseFiles("view/grupos/gruposn.html"))

	ModelosCategorias  = template.Must(template.ParseFiles("view/categorias/listCategoria.html"))
	ModelosCategoriasC = template.Must(template.ParseFiles("view/categorias/AltCategoria.html"))
	ModelosCategoriasN = template.Must(template.ParseFiles("view/categorias/newCategoria.html"))
)
