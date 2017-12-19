package main

import (
	"fmt"
	"net/http"
	"github.com/graphql-go/graphql"
	"encoding/json"
	_ "github.com/lib/pq"
	"callisto/models"
	"callisto/mutations"
	"callisto/queries"
)

type QueryStruct struct {
	Query string `json:"query"`
	OperationName string `json:"operationName"`
	Variables     map[string]interface{} `json:"variables" url:"variables" schema:"variables"`
}

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createProject": mutations.CreateProject,
		"createUser": mutations.CreateUser,
		"createRelease": mutations.CreateRelease,
		"createTask": mutations.CreateTask,
	},
})

// root query
// we just define a trivial example here, since root query is required.
// Test with curl
// curl -g 'http://localhost:8080/graphql?query={lastTodo{id,text,done}}'
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"projectList": queries.ListProjects,
		"project": queries.GetProject,
		"projectByName": queries.GetProjectByName,
		"userList": queries.ListUsers,
		"releaseList": queries.ListReleases,
		"taskList": queries.ListTasks,
	},
})


// define schema, with our rootQuery and rootMutation
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})


func executeQuery(query QueryStruct, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query.Query,
		VariableValues: query.Variables,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	models.InitDB("user=callisto dbname=callisto password=postgrespassword host=172.19.0.2 port=5432 sslmode=disable")
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Body != nil && r.Method == "POST" {
			decoder := json.NewDecoder(r.Body)
			var q QueryStruct
			err := decoder.Decode(&q)
			if err != nil {
				fmt.Printf("%s", err)
			}

			// Debug received query
			s, _ := json.Marshal(q)
			fmt.Println(string(s))

			defer r.Body.Close()
			result := executeQuery(q, schema)
			json.NewEncoder(w).Encode(result)
		}
	})
	// Display some basic instructions
	fmt.Println("Now server is running on port 8080")
	fmt.Println("Access the web app via browser at 'http://localhost:8080'")

	http.ListenAndServe(":8080", nil)
}