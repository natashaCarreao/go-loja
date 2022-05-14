package models

import "github.com/natashacastello/go-loja/db"

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func BuscarTodosOsProdutos() []Produto {

	db := db.ConectaBanco()
	selectAll, err := db.Query("Select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectAll.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectAll.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Id = id
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func CriarNovoProduto(nome, descricao string, quantidade int, preco float64) {
	db := db.ConectaBanco()

	insert, err := db.Prepare("INSERT into produtos (nome, descricao, quantidade, preco) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}
