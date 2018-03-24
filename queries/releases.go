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