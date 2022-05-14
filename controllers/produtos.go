package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/natashacastello/go-loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	produtos := models.BuscarTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)

}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		quatidade, err := strconv.Atoi(r.FormValue("quatidade"))
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)

		if err != nil {
			log.Println("Erro ao convertes string de produtos em inteiro ou float")
		}

		models.CriarNovoProduto(nome, descricao, quatidade, preco)

		http.Redirect(w, r, "/", 301)

	}
}
