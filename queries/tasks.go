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
