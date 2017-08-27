package models

import (
	"log"
	"github.com/graphql-go/graphql"
)

type Project struct {
	Id int `json:"id",db:"id"`
	Name string `json:"name",db:"name"`
	Description string `json:"description",db:"description"`
	Repository string `json:"repository",db:"repository"`
	Url string `json:"url",db:"url"`
	Status int `json:"status",db:"status"`
}


// define custom GraphQL ObjectType `ProjectType` for our Golang struct `Project`
// Note that
// - the fields in our ProjectType maps with the json tags for the fields in our struct
// - the field type matches the field type in our struct
var ProjectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Project",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"repository": &graphql.Field{
			Type: graphql.String,
		},
		"url": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.Int,
		},
	},
})


func AllProjects() ([]*Project, error) {
	rows, err := db.Query(`SELECT * FROM projects`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prjs := make([]*Project, 0)

	for rows.Next() {
		prj := new(Project)
		err := rows.Scan(&prj.Name, &prj.Description, &prj.Repository, &prj.Url, &prj.Status, &prj.Id)
		if err != nil {
			return nil, err
		}
		prjs = append(prjs, prj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return prjs, nil
}

func NewProject(prj Project) (error) {
	stmt, err := db.Prepare("INSERT INTO projects(name, description, repository, url, status) VALUES($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(prj.Name, prj.Description, prj.Repository, prj.Url, prj.Status)
	if err != nil {
		return err
	}
	log.Printf("Result = %+v", res)
	return nil
}