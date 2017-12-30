package models

import (
	"log"
	"github.com/graphql-go/graphql"
)

type Release struct {
	Id int `json:"id",db:"id"`
	ProjectId int `json:"projectId",db:"projectId"`
	Version string `json:"version",db:"version"`
}


// define custom GraphQL ObjectType `ReleaseType` for our Golang struct `Release`
// Note that
// - the fields in our ReleaseType maps with the json tags for the fields in our struct
// - the field type matches the field type in our struct
var ReleaseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Release",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"version": &graphql.Field{
			Type: graphql.String,
		},
	},
})

func AllReleases() ([]*Release, error) {
	rows, err := db.Query(`SELECT * FROM releases`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rels := make([]*Release, 0)

	for rows.Next() {
		rel := new(Release)
		err := rows.Scan(&rel.Id, &rel.ProjectId, &rel.Version)
		if err != nil {
			return nil, err
		}
		rels = append(rels, rel)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rels, nil
}


func ReleasesForProject(id int) ([]*Release, error) {
	rows, err := db.Query(`SELECT * FROM releases WHERE projectId=$1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rels := make([]*Release, 0)

	for rows.Next() {
		rel := new(Release)
		err := rows.Scan(&rel.Id, &rel.ProjectId, &rel.Version)
		if err != nil {
			return nil, err
		}
		rels = append(rels, rel)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rels, nil
}


func NewRelease(rel Release) (error) {
	stmt, err := db.Prepare("INSERT INTO releases(version) VALUES($1)")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(rel.Version)
	if err != nil {
		return err
	}
	log.Printf("Result = %+v", res)
	return nil
}