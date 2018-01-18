package main

import (
	"fmt"

	"github.com/DiegoSantosWS/cms/cone"
	"github.com/DiegoSantosWS/cms/rotas"
)

func init() {
	fmt.Println("Iniciando o servidor")
}

func main() {
	err := cone.Connection()
	if err != nil {
		fmt.Println("Erro ao abrir banco de dandos: ", err.Error())
		return
	}
	rotas.Routers()
}
