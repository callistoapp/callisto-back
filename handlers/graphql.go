package handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"callisto/queries"
	"callisto/mutations"
	gorillaContext "github.com/gorilla/context"
	"context"
)

type QueryStruct struct {
	Query string `json:"query"`
	OperationName string `json:"operationName"`
	Variables     map[string]interface{} `json:"variables" url:"variables" schema:"variables"`
}

// define schema, with our rootQuery and rootMutation
var GraphqlSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queries.Queries,
	Mutation: mutations.Mutations,
})

func ExecuteQuery(query QueryStruct, schema graphql.Schema, ctx context.Context) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query.Query,
		VariableValues: query.Variables,
		Context: ctx,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func GraphqlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://app.callisto.com")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method != "POST" {
		return
	}

	loggedUser := gorillaContext.Get(r, "loggedUser")
	ctx := context.WithValue(context.Background(), "loggedUser", loggedUser)


	if r.Body != nil {
		decoder := json.NewDecoder(r.Body)
		var q QueryStruct
		err := decoder.Decode(&q)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		defer r.Body.Close()
		result := ExecuteQuery(q, GraphqlSchema, ctx)
		json.NewEncoder(w).Encode(result)
		return
	}
}
