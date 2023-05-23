package controllers // vê qual req está vindo e pede para o model buscar os dados no banco de dados

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/MLCavalcante/loja/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html")) // carrega todos os templates

func Index(w http.ResponseWriter, r *http.Request) {	
	todosOsProdutos := models.BuscaTodosOsProdutos() // busca todos os produtos
    templates.ExecuteTemplate(w, "Index", todosOsProdutos)
	
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request){
     if r.Method == "POST" { 
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}
		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}
		models.CriaNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)	
	 }
	 http.Redirect(w, r, "/", 301)

}

func Delete(w http.ResponseWriter, r *http.Request){
	idDoProduto := r.URL.Query().Get("id") // pega o id do produto la na url
	models.DeletaProduto(idDoProduto) // deleta o produto
	http.Redirect(w, r, "/", 301) // redireciona para a pagina inicial
}

func Edit(w http.ResponseWriter, r *http.Request){
	idDoProduto := r.URL.Query().Get("id") // pega o id do produto la na url
	produto := models.EditaProduto(idDoProduto) // edita o produto
	templates.ExecuteTemplate(w, "Edit", produto) // executa o template edit
}

func Update(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		id := r.FormValue("id") // pega o id do produto la na url
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertido, err := strconv.Atoi(id) // converte o id para int
		if err != nil {
			log.Println("Erro na conversão do id para int:", err)
		}

		precoConvertido, err := strconv.ParseFloat(preco, 64) // converte o preço para float
		if err != nil {
			log.Println("Erro na conversão do preço para float64:", err)
		}
		
		quantidadeConvertida, err := strconv.Atoi(quantidade) // converte a quantidade para int
		if err != nil {
			log.Println("Erro na conversão da quantidade para int:", err)
		}
		models.AtualizaProduto(idConvertido, nome, descricao, precoConvertido, quantidadeConvertida) // atualiza o produto
	}
	http.Redirect(w, r, "/", 301) // redireciona para a pagina inicial
}