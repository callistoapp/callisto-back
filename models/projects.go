package models

import (
	"github.com/graphql-go/graphql"
)

type Project struct {
	Id          int        `json:"id",db:"id"`
	Name        string     `json:"name",db:"name"`
	Description string     `json:"description",db:"description"`
	Repository  string     `json:"repository",db:"repository"`
	Url         string     `json:"url",db:"url"`
	Status      int        `json:"status",db:"status"`
	Deleted     int        `json:"deleted",db:"deleted"`
	Tasks       []*Task    `json:"tasks"`
	Releases    []*Release `json:"releases"`
	Statuses    []*Status  `json:"statuses"`
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
		"releases": &graphql.Field{
			Type: graphql.NewList(ReleaseType),
		},
		"statuses": &graphql.Field{
			Type: graphql.NewList(StatusType),
		},
	},
})

func ProjectFromId(id int) (*Project, error) {
	row := db.QueryRow(`SELECT * FROM projects WHERE id=$1 AND deleted = 0`, id)

	prj := new(Project)

	err := row.Scan(&prj.Id, &prj.Name, &prj.Description, &prj.Repository, &prj.Url, &prj.Status, &prj.Deleted)
	if err != nil {
		return nil, err
	}

	prj.Tasks, err = TasksForProject(prj.Id)
	prj.Releases, err = ReleasesForProject(prj.Id)
	prj.Statuses, err = StatusesForProject(prj.Id)

	return prj, nil
}

func ProjectFromName(name string) (*Project, error) {
	row := db.QueryRow(`SELECT * FROM projects WHERE name=$1 AND deleted = 0`, name)

	prj := new(Project)

	err := row.Scan(&prj.Id, &prj.Name, &prj.Description, &prj.Repository, &prj.Url, &prj.Status, &prj.Deleted)
	if err != nil {
		return nil, err
	}

	prj.Tasks, err = TasksForProject(prj.Id)
	prj.Releases, err = ReleasesForProject(prj.Id)
	prj.Statuses, err = StatusesForProject(prj.Id)

	return prj, nil
}

func AllProjects() ([]*Project, error) {
	rows, err := db.Query(`SELECT * FROM projects WHERE deleted = 0`)
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
		err := rows.Scan(&prj.Id, &prj.Name, &prj.Description, &prj.Repository, &prj.Url, &prj.Status, &prj.Deleted)
		if err != nil {
			return nil, err
		}
		prj.Tasks, err = TasksForProject(prj.Id)
		prj.Releases, err = ReleasesForProject(prj.Id)
		prj.Statuses, err = StatusesForProject(prj.Id)

		prjs = append(prjs, prj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return prjs, nil
}

func NewProject(prj Project) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO projects(name, description, repository, url, status, deleted) VALUES($1, $2, $3, $4, $5, $6) RETURNING id", &prj.Name, &prj.Description, &prj.Repository, &prj.Url, &prj.Status, &prj.Deleted).Scan(&id)
	if err != nil {
		return 0, err
	}
	for _, status := range prj.Statuses {
		status.ProjectId = id
		NewStatus(*status)
	}
	return id, nil
}

func UpdateProject(prj Project) error {
	stmt, err := db.Prepare("UPDATE projects set name = $1, description = $2, repository = $3, url = $4 WHERE id = $5")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&prj.Name, &prj.Description, &prj.Repository, &prj.Url, &prj.Id)

	return err
}


func DeleteProject(id int) error {
	stmt, err := db.Prepare("UPDATE projects set deleted = 1 WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)

	return err
}
