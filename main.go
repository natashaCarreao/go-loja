package main

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
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
