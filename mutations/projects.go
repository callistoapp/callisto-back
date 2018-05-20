package mutations

import (
	"github.com/graphql-go/graphql"
	"callisto/models"
	"reflect"
)

/*
curl -g 'http://localhost:8080/graphql?query=mutation+_{createProject(name:"Test",description:"Test",repository:"Test",url:"Test.com",status:2){id,name,description}}'
*/
var CreateProject = &graphql.Field{
	Type:        models.ProjectType, // the return type for this field
	Description: "Create new project",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"repository": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"url": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"statuses": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		name, _ := params.Args["name"].(string)
		description, _ := params.Args["description"].(string)
		repo, _ := params.Args["repository"].(string)
		url, _ := params.Args["url"].(string)
		statuses := reflect.ValueOf(params.Args["statuses"])

		allStatuses := make([]*models.Status, 0)

		for i:=0; i<statuses.Len(); i++ {
			s := new(models.Status)

			s.Name = statuses.Index(i).Interface().(string)
			s.Index = i
			s.Deleted = 0

			allStatuses = append(allStatuses, s)
		}

		newProject := models.Project{
			Name:        	name,
			Description: 	description,
			Repository:  	repo,
			Url:         	url,
			Status:      	0,
			Deleted:   		0,
			Statuses: 		allStatuses,
		}

		id, err := models.NewProject(newProject)

		newProject.Id = id

		return newProject, err
	},
}

/*
curl -g 'http://localhost:8080/graphql?query=mutation+_{createProject(name:"Test",description:"Test",repository:"Test",url:"Test.com",status:2){id,name,description}}'
*/
var UpdateProject = &graphql.Field{
	Type:        models.ProjectType, // the return type for this field
	Description: "Update project",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"repository": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"url": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		id, _ := params.Args["id"].(int)
		name, _ := params.Args["name"].(string)
		description, _ := params.Args["description"].(string)
		repo, _ := params.Args["repository"].(string)
		url, _ := params.Args["url"].(string)

		updatedProject := models.Project{
			Id: 			id,
			Name:        	name,
			Description: 	description,
			Repository:  	repo,
			Url:         	url,
			Status:      	0,
			Deleted:   		0,
		}

		err := models.UpdateProject(updatedProject)
		if err != nil {
			return nil, err
		}

		project, err := models.ProjectFromId(id)
		return project, err
	},
}

/*
curl -g 'http://localhost:8080/graphql?query=mutation+_{createProject(name:"Test",description:"Test",repository:"Test",url:"Test.com",status:2){id,name,description}}'
*/
var DeleteProject = &graphql.Field{
	Type:        models.ProjectType, // the return type for this field
	Description: "Delete existing project",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// marshall and cast the argument value
		id, _ := params.Args["id"].(int)

		err := models.DeleteProject(id)
		return nil, err
	},
}
