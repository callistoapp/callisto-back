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
	"github.com/labstack/gommon/log"
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

func verifyToken(client pb.AuthorizeClient, token *pb.CallistoToken) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	logged, err := client.Authorize(ctx, token)
	if err != nil {
		log.Info("%v.Authorize(_) = _, %v: ", client, err)
		return false, err
	}
	return logged.Logged, nil
}

func isLogged(r *http.Request) (bool, error) {
	conn, err := grpc.Dial("auth:50051",  grpc.WithInsecure())
	if err != nil {
		log.Fatal("fail to dial: %v", err)
		return false, err
	}
	defer conn.Close()
	client := pb.NewAuthorizeClient(conn)
	token, err := r.Cookie("connect.sid")
	if err != nil {
		log.Error("Error occurred while searching connect.sid token : %v", err)
		return false, err
	}
	return verifyToken(client, &pb.CallistoToken{Token: token.Value})
}


func main() {
	models.InitDB("user=callisto dbname=callisto password=postgrespassword host=postgres sslmode=disable")
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://app.callisto.com")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		fmt.Println("Hello from /graphql")

		if r.Body != nil && r.Method == "POST" {
			logged, err := isLogged(r)

			if logged == false || err != nil {
				log.Info("Not authorized !")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("401 - Unauthorized!"))
				return
			}
			decoder := json.NewDecoder(r.Body)
			var q QueryStruct
			err = decoder.Decode(&q)
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "localhost")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		fmt.Println("Hello from /")

		json.NewEncoder(w).Encode("Hello From graphql api")
	})

	http.HandleFunc("/logged", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		fmt.Println("Hello from /logged")
		isLogged(r)
	})
	// Display some basic instructions
	fmt.Println("Now server is running on port 8081")
	fmt.Println("Access the web app via browser at 'http://localhost:8081'")

	http.ListenAndServe(":8081", nil)
}
