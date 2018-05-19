package mutations

import (
"github.com/graphql-go/graphql"
"callisto/models"
)

/*
curl -g 'http://localhost:8080/graphql?query=mutation+_{createTask(name:"Task1"){id,text,done}}'
*/
var EditStatus = &graphql.Field{
	Type:        models.StatusType,
	Description: "Edit status information",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		name, _ := params.Args["name"].(string)
		id, _ := params.Args["id"].(int)

		// perform mutation operation here
		// for e.g. create a Task and save to DB.
		updatedStatus := models.Status{
			Name: name,
			Id:   id,
		}

		status, err := models.EditStatus(updatedStatus)

		return status, err
	},
}