package repository

import (
	"fmt"
	"movieInfo/internal/database"
	"movieInfo/internal/models"
)

func SearchMovies(searchString string, page int) ([]models.Movie, error) {
	searchObject := database.PaginatedSearch{
		Table:          "movies",
		QueryString:    searchString,
		QueryField:     "name",
		Page:           page,
		RecordsPerPage: 10,
	}
	results, err := database.MakePaginatedSearch(searchObject)
	if err != nil {
		return nil, err
	}
	movies, errTwo := models.GetMoviesFromRows(results.Records)
	if errTwo != nil {
		return nil, errTwo
	}
	return movies, nil
}

func GetMovieByID(id int) (models.Movie, error) {
	row, err := database.GetById("movies", id)
	if err != nil {
		print("error: ", err)
		return models.Movie{}, err
	}
	movie, errTwo := models.GetMovieFromRow(row)
	if errTwo != nil {
		print("error 2: ", err)
		return models.Movie{}, errTwo

	}
	return movie, nil
}

func GetActorsInAMovie(movieId int) ([]models.Actor, error) {
	queryString :=
		fmt.Sprintf("SELECT a.id, a.name, r.role "+
			"FROM moviesearch.roles as r "+
			"INNER JOIN "+
			"moviesearch.actors as a "+
			"ON a.id = r.actors_id "+
			"WHERE movies_id = %v ", movieId)
	println("queryString: ", queryString)
	rows, err := database.MakeCustomQuery(queryString)
	if err != nil {
		return []models.Actor{}, err
	}
	actors, err := models.GetActorsAndRolesFromRows(rows)
	if err != nil {
		return []models.Actor{}, err
	}
	return actors, nil
}
