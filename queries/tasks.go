package queries

import (
	"github.com/graphql-go/graphql"
	"callisto/models"
)

//curl -g 'http://localhost:8080/graphql?query={taskList{id,name,desription}}'
var ListTasks = &graphql.Field{
	Type:        graphql.NewList(models.TaskType),
	Description: "List of tasks",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		taskList, err := models.AllTasks()
		return taskList, err
	},
}

//curl -g 'http://localhost:8080/graphql?query={taskList{id,name,desription}}'
var GetTask = &graphql.Field{
	Type:        models.TaskType,
	Description: "Find a task by its id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// marshall and cast the argument value
		id, _ := params.Args["id"].(int)

		task, err := models.TaskFromId(id)

		if err != nil {
			return nil, err
		}

		return task, err
	},
}
