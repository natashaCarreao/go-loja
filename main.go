package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/natashacastello/go-loja/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
