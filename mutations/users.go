package mutations

import (
	"github.com/graphql-go/graphql"
	"callisto/models"
)

/*
curl -g 'http://localhost:8080/graphql?query=mutation+_{createUser(text:"My+new+user"){id,text,done}}'
*/
var CreateUser = &graphql.Field{
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
			Name:    name,
			Email:   email,
			Phone:   phone,
			Deleted: 0,
		}

		err := models.NewUser(newUser)

		return newUser, err
	},
}
