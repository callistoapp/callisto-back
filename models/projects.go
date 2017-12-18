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
	Tasks []*Task `json:"tasks"`
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
		"tasks": &graphql.Field{
			Type: graphql.NewList(TaskType),
		},
	},
})


func ProjectFromId(id int) (*Project, error) {
	row := db.QueryRow(`SELECT * FROM projects WHERE id=$1`, id)

	prj := new(Project)

	err := row.Scan(&prj.Name, &prj.Description, &prj.Repository, &prj.Url, &prj.Status, &prj.Id)
	if err != nil {
		return nil, err
	}

	prj.Tasks, err = TasksForProject(prj.Id)

	return prj, nil
}


func ProjectFromName(name string) (*Project, error) {
	row := db.QueryRow(`SELECT * FROM projects WHERE name=$1`, name)

	prj := new(Project)

	err := row.Scan(&prj.Name, &prj.Description, &prj.Repository, &prj.Url, &prj.Status, &prj.Id)
	if err != nil {
		return nil, err
	}

	prj.Tasks, err = TasksForProject(prj.Id)

	return prj, nil
}


func AllProjects() ([]*Project, error) {
	rows, err := db.Query(`SELECT * FROM projects`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prjs := make([]*Project, 0)

	for rows.Next() {
		prj := new(Project)
		if err != nil {
			return nil, err
		}
		err := rows.Scan(&prj.Name, &prj.Description, &prj.Repository, &prj.Url, &prj.Status, &prj.Id)
		if err != nil {
			return nil, err
		}
		prj.Tasks, err = TasksForProject(prj.Id)

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