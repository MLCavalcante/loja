package main

import (
	"net/http"

	"github.com/MLCavalcante/loja/routes"
	_ "github.com/MLCavalcante/loja/routes"
)







func main() {
	routes.CarregaRotas()
   
	http.ListenAndServe(":8000", nil)
}

