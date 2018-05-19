package main

import (
	"net/http"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"callisto/models"
	"callisto/handlers"
	"callisto/middlewares"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/graphql", handlers.GraphqlHandler)
	r.HandleFunc("/", handlers.HomeHandler)
	r.Use(middlewares.AuthMiddleware)

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)

	models.InitDB("user=callisto dbname=callisto password=postgrespassword host=postgres sslmode=disable")

	http.ListenAndServe(":8081", n)
}
