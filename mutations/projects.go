package mutations

import (
	"github.com/graphql-go/graphql"
	"callisto/models"
	"reflect"
)

/*
curl -g 'http://localhost:8080/graphql?query=mutation+_{createProject(name:"Test",description:"Test",repository:"Test",url:"Test.com",status:2){id,name,description}}'
*/
var CreateProject = &graphql.Field{
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
		"statuses": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		name, _ := params.Args["name"].(string)
		description, _ := params.Args["description"].(string)
		repo, _ := params.Args["repo"].(string)
		url, _ := params.Args["url"].(string)
		statuses := reflect.ValueOf(params.Args["statuses"])

		allStatuses := make([]*models.Status, 0)

		for i:=0; i<statuses.Len(); i++ {
			s := new(models.Status)

			s.Name = statuses.Index(i).Interface().(string)
			s.Index = i
			s.Deleted = 0

			allStatuses = append(allStatuses, s)
		}

		newProject := models.Project{
			Name:        	name,
			Description: 	description,
			Repository:  	repo,
			Url:         	url,
			Status:      	0,
			Deleted:   		0,
			Statuses: 		allStatuses,
		}

		id, err := models.NewProject(newProject)

		return id, err
	},
}

///*
//	curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTodo(id:"a",done:true){id,text,done}}'
//*/
//var UpdateProject = &graphql.Field{
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
