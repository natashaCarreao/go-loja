package main

import (
	"html/template"
	"net/http"
	_"github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := conectaBanco()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Camiseta", "Preta", 200, 2},
		{"Fone", "Otimo", 240, 25},
		{"Tenis", "Preto e branco", 100, 12},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func conectaBanco() *sql.DB {
	conexao := "user=postgres dbname=goloja password=postgres host=postgres-db sslmode=desable"
	db, err sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
