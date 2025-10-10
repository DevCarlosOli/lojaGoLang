package routes

import (
	"lojaGoLang/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
