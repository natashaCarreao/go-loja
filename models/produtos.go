package models

import "github.com/go-loja/db"

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

func DeletarPorId(id string) {
	db := db.ConectaBanco()

	delete, err := db.Prepare("DELETE produtos WHERE id = $1")

	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)

	defer db.Close()
}

func BuscarProdutoPorId(id string) Produto {
	db := db.ConectaBanco()

	findById, err := db.Query("SELECT * FROM produtos where = $1 ORDER BY id", id)

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	for findById.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = findById.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Id = id
		p.Preco = preco
		p.Quantidade = quantidade
	}

	defer db.Close()

	return p

}

func AtualizarProduto(id, nome, descricao string, quantidade int, preco float64) {
	db := db.ConectaBanco()

	update, err := db.Prepare("UPDATE produtos SET nome = $1, descricao = $2, quantidade =$3, preco $4 WHERE id = $5")

	if err != nil {
		panic(err.Error())
	}

	update.Exec(nome, descricao, quantidade, preco, id)

	defer db.Close()
}
