package mutations

import (
	"github.com/graphql-go/graphql"
	"callisto/models"
)

/*
curl -g 'http://localhost:8080/graphql?query=mutation+_{createTask(name:"Task1"){id,text,done}}'
*/
var CreateTask = &graphql.Field{
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
			Name:        name,
			Description: description,
			Type:        taskType,
			Status:      status,
		}

		err := models.NewTask(newTask)

		return newTask, err
	},
}
