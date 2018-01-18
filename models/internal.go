package models

import (
	"fmt"
	"net/http"

	ctrl "github.com/DiegoSantosWS/cms/controller"
)

//Internal abre arquivo interno
func Internal(w http.ResponseWriter, r *http.Request) {
	//h.CheckSession(w, r)
	if err := ctrl.ModelosInternal.ExecuteTemplate(w, "interno.html", nil); err != nil {
		http.Error(w, "[TEMPLATE LOGIN], erro no carregamento", http.StatusInternalServerError)
		fmt.Println("LOGIN: erro na execução do modelo", err.Error())
	}
}
