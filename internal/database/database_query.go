package database

import (
	"database/sql"
	"fmt"
)

// / Create a type for paginated results
type PaginatedResults struct {
	TotalRecords int
	TotalPages   int
	CurrentPage  int
	Records      *sql.Rows
}

type PaginatedSearch struct {
	Table          string
	QueryString    string
	QueryField     string
	Page           int
	RecordsPerPage int
}

func MakeCustomQuery(queryString string) (*sql.Rows, error) {
	/// Get a database connection
	db := GetConnection()
	rows, err := db.Query(queryString)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func GetById(table string, id int) (*sql.Row, error) {
	db := GetConnection()
	queryString := fmt.Sprintf("SELECT * FROM %v WHERE id = %v", table,id)
	println("queryString: ", queryString)
	row := db.QueryRow(queryString)

	return row, nil
}

func MakePaginatedSearch(searchObject PaginatedSearch) (PaginatedResults, error) {
	/// Get a database connection
	db := GetConnection()
	queryString := fmt.Sprintf("SELECT * FROM %v WHERE %v LIKE '%%%v%%' LIMIT %v OFFSET %v", searchObject.Table, searchObject.QueryField, searchObject.QueryString, searchObject.RecordsPerPage, searchObject.Page)
	countString := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE %v LIKE '%%%v%%'", searchObject.Table, searchObject.QueryField, searchObject.QueryString)
	countRow := db.QueryRow(countString)
	var count int
	countRow.Scan(&count)

	println("count: ", count)
	println("queryString: ", queryString)

	rows, err := db.Query(queryString)
	if err != nil {
		return PaginatedResults{}, err
	}
	//defer rows.Close()
	println("rows: ", rows)
	return PaginatedResults{
		TotalRecords: count,
		TotalPages:   count / searchObject.RecordsPerPage,
		CurrentPage:  searchObject.Page,
		Records:      rows,
	}, nil

}
