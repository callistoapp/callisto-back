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
		"projectId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
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
		projectId, _ := params.Args["projectId"].(int)
		description, _ := params.Args["description"].(string)
		taskType, _ := params.Args["type"].(int)
		statusId, _ := params.Args["status"].(int)

		// perform mutation operation here
		// for e.g. create a Task and save to DB.
		newTask := models.Task{
			Name:        name,
			ProjectId:   projectId,
			Description: description,
			Type:        taskType,
			StatusId:    statusId,
			Deleted:     0,
		}

		err := models.NewTask(newTask)

		return newTask, err
	},
}

/*
curl -g 'http://localhost:8080/graphql?query=mutation+_{createTask(name:"Task1"){id,text,done}}'
*/
var MoveTask = &graphql.Field{
	Type:        graphql.NewList(models.TaskType), // the return type for this field
	Description: "Create new task",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"statusId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		id, _ := params.Args["id"].(int)
		statusId, _ := params.Args["statusId"].(int)

		err := models.MoveTask(id, statusId)

		if err != nil {
			return nil, err
		}

		allTasks, err := models.AllTasks()

		return allTasks, err
	},
}

/*
curl -g 'http://localhost:8080/graphql?query=mutation+_{createTask(name:"Task1"){id,text,done}}'
*/
var EditTask = &graphql.Field{
	Type:        graphql.NewList(models.TaskType), // the return type for this field
	Description: "Create new task",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
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
		id, _ := params.Args["id"].(int)
		description, _ := params.Args["description"].(string)
		taskType, _ := params.Args["type"].(int)
		statusId, _ := params.Args["status"].(int)

		// perform mutation operation here
		// for e.g. create a Task and save to DB.
		newTask := models.Task{
			Name:        name,
			Id:   		 id,
			Description: description,
			Type:        taskType,
			StatusId:    statusId,
			Deleted:     0,
		}

		err := models.EditTask(newTask)

		if err != nil {
			return nil, err
		}

		allTasks, err := models.AllTasks()

		return allTasks, err
	},
}

/*
curl -g 'http://localhost:8080/graphql?query=mutation+_{createTask(name:"Task1"){id,text,done}}'
*/
var DeleteTask = &graphql.Field{
	Type:        graphql.NewNonNull(graphql.Int), // the return type for this field
	Description: "Delete a task",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		id, _ := params.Args["id"].(int)

		err := models.DeleteTask(id)

		if err != nil {
			return 0, err
		}

		return 1, err
	},
}
