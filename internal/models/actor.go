package models

import (
	"database/sql"
	"encoding/json"
)

type Actor struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func GetActorsFromRows(rows *sql.Rows) ([]Actor, error) {
	actors := []Actor{}
	for rows.Next() {
		var actor Actor
		err := rows.Scan(&actor.ID, &actor.Name)
		if err != nil {
			return nil, err
		}
		actors = append(actors, actor)
	}
	return actors, nil
}

func GetActorsAndRolesFromRows(rows *sql.Rows) ([]Actor, error) {
	actors := []Actor{}
	for rows.Next() {
		var actor Actor
		err := rows.Scan(&actor.ID, &actor.Name, &actor.Role)
		if err != nil {
			return nil, err
		}
		actors = append(actors, actor)
	}
	return actors, nil
}

func GetActorFromRow(row *sql.Row) (Actor, error) {
	var actor Actor
	err := row.Scan(&actor.ID, &actor.Name)
	if err != nil {
		return Actor{}, err
	}
	return actor, nil
}

func ToJsonActor(actors []Actor) ([]byte, error) {
	result, err := json.Marshal(actors)
	if err != nil {
		return nil, err
	}
	return result, nil
}
