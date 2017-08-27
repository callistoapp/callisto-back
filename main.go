package main

import (
	"fmt"
	"net/http"
	"github.com/graphql-go/graphql"
	"encoding/json"
	_ "github.com/lib/pq"
	"callisto/models"
)

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		/*
			curl -g 'http://localhost:8080/graphql?query=mutation+_{createProject(name:"Test",description:"Test",repository:"Test",url:"Test.com",status:2){id,name,description}}'
		*/
		"createProject": &graphql.Field{
			Type:        models.ProjectType, // the return type for this field
			Description: "Create new project",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"repository": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"url": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"status": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				// marshall and cast the argument value
				name, _ := params.Args["name"].(string)
				description, _ := params.Args["description"].(string)
				repo, _ := params.Args["repo"].(string)
				url, _ := params.Args["url"].(string)

				// perform mutation operation here
				// for e.g. create a Project and save to DB.
				newProject := models.Project{
					Name:   name,
					Description: description,
					Repository: repo,
					Url: url,
					Status: 0,
				}

				err := models.NewProject(newProject)

				return newProject, err
			},
		},
		/*
			curl -g 'http://localhost:8080/graphql?query=mutation+_{createUser(text:"My+new+todo"){id,text,done}}'
		*/
		"createUser": &graphql.Field{
			Type:        models.UserType, // the return type for this field
			Description: "Create new user",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"phone": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				// marshall and cast the argument value
				name, _ := params.Args["name"].(string)
				email, _ := params.Args["email"].(string)
				phone, _ := params.Args["phone"].(string)

				// perform mutation operation here
				// for e.g. create a User and save to DB.
				newUser := models.User{
					Name:   name,
					Email: email,
					Phone: phone,
				}

				err := models.NewUser(newUser)

				return newUser, err
			},
		},
		/*
			curl -g 'http://localhost:8080/graphql?query=mutation+_{createRelease(version:"v1.0.0"){id,version}}'
		*/
		"createRelease": &graphql.Field{
			Type:        models.ReleaseType, // the return type for this field
			Description: "Create new release",
			Args: graphql.FieldConfigArgument{
				"version": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				// marshall and cast the argument value
				version, _ := params.Args["version"].(string)

				// perform mutation operation here
				// for e.g. create a Release and save to DB.
				newRelease := models.Release{
					Version:   version,
				}

				err := models.NewRelease(newRelease)

				return newRelease, err
			},
		},
		/*
			curl -g 'http://localhost:8080/graphql?query=mutation+_{createTask(name:"Task1"){id,text,done}}'
		*/
		"createTask": &graphql.Field{
			Type:        models.TaskType, // the return type for this field
			Description: "Create new task",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"type": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"status": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				// marshall and cast the argument value
				name, _ := params.Args["name"].(string)
				description, _ := params.Args["description"].(string)
				taskType, _ := params.Args["type"].(int)
				status, _ := params.Args["status"].(int)

				// perform mutation operation here
				// for e.g. create a Task and save to DB.
				newTask := models.Task{
					Name:   name,
					Description: description,
					Type: taskType,
					Status: status,
				}

				err := models.NewTask(newTask)

				return newTask, err
			},
		},
		///*
		//	curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTodo(id:"a",done:true){id,text,done}}'
		//*/
		//"updateTodo": &graphql.Field{
		//	Type:        todoType, // the return type for this field
		//	Description: "Update existing todo, mark it done or not done",
		//	Args: graphql.FieldConfigArgument{
		//		"done": &graphql.ArgumentConfig{
		//			Type: graphql.Boolean,
		//		},
		//		"id": &graphql.ArgumentConfig{
		//			Type: graphql.NewNonNull(graphql.String),
		//		},
		//	},
		//	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//		// marshall and cast the argument value
		//		done, _ := params.Args["done"].(bool)
		//		id, _ := params.Args["id"].(string)
		//		affectedTodo := Todo{}
		//
		//		// Search list for todo with id and change the done variable
		//		for i := 0; i < len(TodoList); i++ {
		//			if TodoList[i].ID == id {
		//				TodoList[i].Done = done
		//				// Assign updated todo so we can return it
		//				affectedTodo = TodoList[i]
		//				break
		//			}
		//		}
		//		// Return affected todo
		//		return affectedTodo, nil
		//	},
		//},
	},
})

// root query
// we just define a trivial example here, since root query is required.
// Test with curl
// curl -g 'http://localhost:8080/graphql?query={lastTodo{id,text,done}}'
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{

		///*
		//   curl -g 'http://localhost:8080/graphql?query={todo(id:"b"){id,text,done}}'
		//*/
		//"todo": &graphql.Field{
		//	Type:        todoType,
		//	Description: "Get single todo",
		//	Args: graphql.FieldConfigArgument{
		//		"id": &graphql.ArgumentConfig{
		//			Type: graphql.String,
		//		},
		//	},
		//	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//
		//		idQuery, isOK := params.Args["id"].(string)
		//		if isOK {
		//			// Search for el with id
		//			for _, todo := range TodoList {
		//				if todo.ID == idQuery {
		//					return todo, nil
		//				}
		//			}
		//		}
		//
		//		return Todo{}, nil
		//	},
		//},


		/*
		   curl -g 'http://localhost:8080/graphql?query={projectList{id,name,desription}}'
		*/
		"projectList": &graphql.Field{
			Type:        graphql.NewList(models.ProjectType),
			Description: "List of projects",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				projectList, err := models.AllProjects()
				return projectList, err
			},
		},

		/*
		   curl -g 'http://localhost:8080/graphql?query={userList{id,name,desription}}'
		*/
		"userList": &graphql.Field{
			Type:        graphql.NewList(models.UserType),
			Description: "List of users",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				userList, err := models.AllUsers()
				return userList, err
			},
		},

		/*
		   curl -g 'http://localhost:8080/graphql?query={releaseList{id,name,desription}}'
		*/
		"releaseList": &graphql.Field{
			Type:        graphql.NewList(models.ReleaseType),
			Description: "List of releases",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				releaseList, err := models.AllReleases()
				return releaseList, err
			},
		},

		/*
		   curl -g 'http://localhost:8080/graphql?query={taskList{id,name,desription}}'
		*/
		"taskList": &graphql.Field{
			Type:        graphql.NewList(models.TaskType),
			Description: "List of tasks",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				taskList, err := models.AllTasks()
				return taskList, err
			},
		},

		//"lastTodo": &graphql.Field{
		//	Type:        todoType,
		//	Description: "Last todo added",
		//	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//		return TodoList[len(TodoList)-1], nil
		//	},
		//},

	},
})


// define schema, with our rootQuery and rootMutation
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})


func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	models.InitDB("user=postgres dbname=callisto password=291092 host=localhost port=5432 sslmode=disable")
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	// Display some basic instructions
	fmt.Println("Now server is running on port 8080")
	fmt.Println("Access the web app via browser at 'http://localhost:8080'")

	http.ListenAndServe(":8080", nil)
}