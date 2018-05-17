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
	r := mux.NewRouter().StrictSlash(true)
	graphqlRouter := r.PathPrefix("/graphql").Subrouter()
	graphqlRouter.HandleFunc("/", handlers.GraphqlHandler)
	graphqlRouter.Use(middlewares.AuthMiddleware)

	r.HandleFunc("/", handlers.HomeHandler)

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)

	models.InitDB("user=callisto dbname=callisto password=postgrespassword host=postgres sslmode=disable")

	http.ListenAndServe(":8081", n)
}
