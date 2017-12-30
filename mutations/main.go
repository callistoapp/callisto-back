package mutations

import (
	"github.com/graphql-go/graphql"
)

// root mutation
var Mutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createProject": CreateProject,
		"createUser": CreateUser,
		"createRelease": CreateRelease,
		"createTask": CreateTask,
		"moveTask": MoveTask,
	},
})