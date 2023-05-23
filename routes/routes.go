package routes

import (
	"net/http"
	"github.com/MLCavalcante/loja/controllers"
	
)

func CarregaRotas() {
	 http.HandleFunc("/", controllers.Index)// sempre que tiver uma req para / quem atende é o index
	 http.HandleFunc("/new", controllers.New)
	 http.HandleFunc("/insert", controllers.Insert)
	 http.HandleFunc("/delete", controllers.Delete)
	 http.HandleFunc("/edit", controllers.Edit)
	 http.HandleFunc("/update", controllers.Update)
}