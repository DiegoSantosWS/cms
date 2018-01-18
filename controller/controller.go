package controller

import (
	"html/template"
)

var (
	ModelosLogin        = template.Must(template.ParseFiles("view/login.html"))
	ModelosCadastrar    = template.Must(template.ParseFiles("view/cad-user.html"))
	ModelosInternal     = template.Must(template.ParseFiles("view/interno.html"))
	ModelosUsuarioslist = template.Must(template.ParseFiles("view/usuarios/listUser.html"))
)
