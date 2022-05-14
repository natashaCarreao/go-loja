package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/go-loja/models"
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

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	models.DeletarPorId(id)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	produto := models.BuscarProdutoPorId(id)

	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		id := r.FormValue("id")
		quatidade, err := strconv.Atoi(r.FormValue("quatidade"))
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)

		if err != nil {
			log.Println("Erro ao convertes string de produtos em inteiro ou float")
		}

		models.AtualizarProduto(id, nome, descricao, quatidade, preco)

		http.Redirect(w, r, "/", 301)
	}
}
