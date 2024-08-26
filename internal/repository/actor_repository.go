package repository

import (
	"movieInfo/internal/database"
	"movieInfo/internal/models"
)

func SearchActors(searchString string, page int) ([]models.Actor, error) {
	searchObject := database.PaginatedSearch{
		Table:          "actors",
		QueryString:    searchString,
		QueryField:     "name",
		Page:           page,
		RecordsPerPage: 10,
	}
	results, err := database.MakePaginatedSearch(searchObject)
	if err != nil {
		return nil, err
	}
	actors, err := models.GetActorsFromRows(results.Records)
	if err != nil {
		return nil, err
	}
	return actors, err
}

func GetActorByID(id int) (models.Actor, error) {
	row, err := database.GetById("actors", id)
	if err != nil {
		print("error", err)
		return models.Actor{}, err
	}
	actor, errTwo := models.GetActorFromRow(row)
	if errTwo != nil {
		print("error 2 ", err)
		return models.Actor{}, errTwo
	}
	return actor, nil
}
