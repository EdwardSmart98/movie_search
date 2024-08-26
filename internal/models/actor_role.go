package models

import "database/sql"

type ActorRole struct {
	ActorID  int    `json:"actors_id"`
	MovieID  int    `json:"movies_id"`
	RoleName string `json:"role"`
}

func GetRolesFromRows(rows *sql.Rows) ([]ActorRole, error) {
	roles := []ActorRole{}
	for rows.Next() {
		var role ActorRole
		err := rows.Scan(&role.MovieID, &role.RoleName, &role.ActorID)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}
