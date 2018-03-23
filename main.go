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
	"google.golang.org/grpc"
	pb "callisto/authorization"
	"time"
	"golang.org/x/net/context"
)


type QueryStruct struct {
	Query string `json:"query"`
	OperationName string `json:"operationName"`
	Variables     map[string]interface{} `json:"variables" url:"variables" schema:"variables"`
}

// define schema, with our rootQuery and rootMutation
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queries.Queries,
	Mutation: mutations.Mutations,
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

func verifyToken(client pb.AuthorizeClient, token *pb.CallistoToken) {
	fmt.Printf("Verifying token: (%d)\n", token.Token)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	logged, err := client.Authorize(ctx, token)
	if err != nil {
		fmt.Printf("%v.Authorize(_) = _, %v: ", client, err)
	}
	fmt.Println("Looks like a success ???")
	fmt.Println(logged.Logged)
}


func main() {
	models.InitDB("user=callisto dbname=callisto password=postgrespassword host=postgres sslmode=disable")
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		fmt.Println("Hello from /graphql")

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
			return
		}
		json.NewEncoder(w).Encode("please use post")
	})

	http.HandleFunc("/logged", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		fmt.Println("Hello from /logged")
		conn, err := grpc.Dial("auth:50051",  grpc.WithInsecure())
		if err != nil {
			fmt.Printf("fail to dial: %v", err)
		}
		defer conn.Close()
		client := pb.NewAuthorizeClient(conn)
		verifyToken(client, &pb.CallistoToken{Token: "nifejnovervrenvrereoj"})
	})
	// Display some basic instructions
	fmt.Println("Now server is running on port 8081")
	fmt.Println("Access the web app via browser at 'http://localhost:8081'")

	http.ListenAndServe(":8081", nil)
}
