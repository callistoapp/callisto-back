package queries

import (
	"github.com/graphql-go/graphql"
	"callisto/models"
)

// curl -g 'http://localhost:8080/graphql?query={releaseList{id,name,desription}}'
var ListReleases = &graphql.Field{
	Type:        graphql.NewList(models.ReleaseType),
	Description: "List of releases",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		releaseList, err := models.AllReleases()
		return releaseList, err
	},
}

// curl -g 'http://localhost:8080/graphql?query={releaseList{id,name,desription}}'
var GetReleasesForProject = &graphql.Field{
	Type:        graphql.NewList(models.ReleaseType),
	Description: "List of releases for a project",
	Args: graphql.FieldConfigArgument{
		"projectId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		projectId, _ := params.Args["projectId"].(int)

		releaseList, err := models.ReleasesForProject(projectId)
		return releaseList, err
	},
}