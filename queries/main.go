package queries

import (
	"github.com/graphql-go/graphql"
)

// root query
// we just define a trivial example here, since root query is required.
// Test with curl
// curl -g 'http://localhost:8080/graphql?query={lastTodo{id,text,done}}'
var Queries = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"projectList": ListProjects,
		"project": GetProject,
		"projectByName": GetProjectByName,
		"userList": ListUsers,
		"releaseList": ListReleases,
		"taskList": ListTasks,
	},
})