package models

import (
	"github.com/graphql-go/graphql"
)

type Status struct {
	Id          int    `json:"id",db:"id"`
	ProjectId   int    `json:"projectId",db:"projectId"`
	Name        string `json:"name",db:"name"`
	Description string `json:"description",db:"description"`
	Index       int    `json:"type",db:"type"`
}

// define custom GraphQL ObjectType `TaskType` for our Golang struct `Task`
// Note that
// - the fields in our TaskType maps with the json tags for the fields in our struct
// - the field type matches the field type in our struct
var StatusType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Status",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"projectId": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"index": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

func AllStatuses() ([]*Status, error) {
	rows, err := db.Query(`SELECT * FROM statuses`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stts := make([]*Status, 0)

	for rows.Next() {
		stt := new(Status)
		err := rows.Scan(&stt.Id, &stt.ProjectId, &stt.Name, &stt.Description, &stt.Index)
		if err != nil {
			return nil, err
		}
		stts = append(stts, stt)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return stts, nil
}

func StatusesForProject(id int) ([]*Status, error) {
	rows, err := db.Query(`SELECT * FROM statuses WHERE projectId=$1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stts := make([]*Status, 0)

	for rows.Next() {
		stt := new(Status)
		err := rows.Scan(&stt.Id, &stt.ProjectId, &stt.Name, &stt.Description, &stt.Index)
		if err != nil {
			return nil, err
		}
		stts = append(stts, stt)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return stts, nil
}

func NewStatus(stt Status) (error) {
	stmt, err := db.Prepare("INSERT INTO statuses(projectId, name, description, index) VALUES($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(stt.ProjectId, stt.Name, stt.Description, stt.Index)
	if err != nil {
		return err
	}
	return nil
}
