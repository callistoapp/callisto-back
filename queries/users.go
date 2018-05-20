package queries

import (
	"github.com/graphql-go/graphql"
	"callisto/models"
)

//curl -g 'http://localhost:8080/graphql?query={userList{id,name,desription}}'
var ListUsers = &graphql.Field{
	Type:        graphql.NewList(models.UserType),
	Description: "List of users",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		userList, err := models.AllUsers()
		return userList, err
	},
}

//curl -g 'http://localhost:8080/graphql?query={userList{id,name,desription}}'
var GetLoggedUser = &graphql.Field{
	Type:        models.LoggedUserType,
	Description: "Logged user",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		loggedUser := p.Context.Value("loggedUser").(models.AuthenticatedUser)
		return loggedUser, nil
	},
}