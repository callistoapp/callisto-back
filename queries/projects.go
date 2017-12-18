package queries

import (
	"github.com/graphql-go/graphql"
	"callisto/models"
	"github.com/go-siris/siris/core/errors"
)

//curl -g 'http://localhost:8080/graphql?query={projectList{id,name,desription}}'
var ListProjects = &graphql.Field{
	Type:        graphql.NewList(models.ProjectType),
	Description: "List of projects",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		projectList, err := models.AllProjects()
		return projectList, err
	},
}

//curl -g 'http://localhost:8080/graphql?query={projectList{id,name,desription}}'
var GetProject = &graphql.Field{
	Type:        models.ProjectType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Description: "One project",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(int)
		if isOK {
			projectList, err := models.ProjectFromId(idQuery)
			return projectList, err
		}
		return models.Project{}, errors.New("Id is not ok")
	},
}

//curl -g 'http://localhost:8080/graphql?query={projectByName{id,name,desription}}'
var GetProjectByName = &graphql.Field{
	Type:        models.ProjectType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Description: "One project",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name, isOK := p.Args["name"].(string)
		if isOK {
			project, err := models.ProjectFromName(name)
			return project, err
		}
		return models.Project{}, errors.New("Name is not ok")
	},
}