package models

import (
	"log"
	"github.com/graphql-go/graphql"
)

type Task struct {
	Id int `json:"id",db:"id"`
	Name string `json:"name",db:"name"`
	Description string `json:"description",db:"description"`
	Type int `json:"type",db:"type"`
	Status int `json:"status",db:"status"`
}

// define custom GraphQL ObjectType `TaskType` for our Golang struct `Task`
// Note that
// - the fields in our TaskType maps with the json tags for the fields in our struct
// - the field type matches the field type in our struct
var TaskType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Task",
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
		"type": &graphql.Field{
			Type: graphql.Int,
		},
		"status": &graphql.Field{
			Type: graphql.Int,
		},
	},
})


func AllTasks() ([]*Task, error) {
	rows, err := db.Query(`SELECT * FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tsks := make([]*Task, 0)

	for rows.Next() {
		tsk := new(Task)
		err := rows.Scan(&tsk.Name, &tsk.Description, &tsk.Type, &tsk.Status, &tsk.Id)
		if err != nil {
			return nil, err
		}
		tsks = append(tsks, tsk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tsks, nil
}

func NewTask(tsk Task) (error) {
	stmt, err := db.Prepare("INSERT INTO tasks(name, description, type, status) VALUES($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(tsk.Name, tsk.Description, tsk.Type, tsk.Status)
	if err != nil {
		return err
	}
	log.Printf("Result = %+v", res)
	return nil
}