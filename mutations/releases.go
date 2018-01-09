package mutations

import (
	"github.com/graphql-go/graphql"
	"callisto/models"
)

/*
curl -g 'http://localhost:8080/graphql?query=mutation+_{createRelease(version:"v1.0.0"){id,version}}'
*/
var CreateRelease = &graphql.Field{
	Type:        models.ReleaseType, // the return type for this field
	Description: "Create new release",
	Args: graphql.FieldConfigArgument{
		"version": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		version, _ := params.Args["version"].(string)

		// perform mutation operation here
		// for e.g. create a Release and save to DB.
		newRelease := models.Release{
			Version: version,
			Deleted: 0,
		}

		err := models.NewRelease(newRelease)

		return newRelease, err
	},
}