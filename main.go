package main

import (
	"lojaGoLang/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
