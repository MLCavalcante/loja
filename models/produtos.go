package models

import(
	 "github.com/MLCavalcante/loja/db"
)

type Produto struct {
	Id int
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto{
	db := db.ConectaComBancoDeDados() // conecta com o banco de dados

	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id asc") // faz a query no banco de dados
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}
	for selectDeTodosOsProdutos.Next() { // percorre todos os produtos
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade) // pega os valores do banco de dados
		if err != nil {
			panic(err.Error())
		}
		p.Id = id // adiciona os valores no produto
		p.Nome = nome 
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		produtos = append(produtos, p) // adiciona o produto no array de produtos

	}
	defer db.Close() // fecha a conexão com o banco de dados
	return produtos 

}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db:= db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)") // prepara a query
	if err != nil {
		panic(err.Error())
	}
	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade) // executa a query
	defer db.Close()

}

func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()

	deletarOProduto, err := db.Prepare("delete from produtos where id=$1") // prepara a query
	if err != nil {
		panic(err.Error())
	}
	deletarOProduto.Exec(id) // executa a query
	defer db.Close()
}

func EditaProduto(id string) Produto { // retorna um produto
	db := db.ConectaComBancoDeDados()

	produtoDoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) { 
     db := db.ConectaComBancoDeDados()

	 atualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5") // prepara a query
	 if err != nil {
		 panic(err.Error())
	 }
	 atualizaProduto.Exec(nome, descricao, preco, quantidade, id) // executa a query
	 defer db.Close()
} 
