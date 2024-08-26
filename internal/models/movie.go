package models

import (
	"database/sql"
	"encoding/json"
)

type Movie struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Date        int     `json:"date"`
	Tagline     string  `json:"tagline"`
	Description string  `json:"description"`
	Minute      int     `json:"minute"`
	Rating      float32 `json:"rating"`
}

func GetMovieFromRow(row *sql.Row) (Movie, error) {
	var movie Movie
	err := row.Scan(&movie.Id, &movie.Name, &movie.Date, &movie.Tagline, &movie.Description, &movie.Minute, &movie.Rating)
	if err != nil {
		print("error failed to scan row: ", err.Error())
		return Movie{}, err
	}
	return movie, nil
}

func GetMoviesFromRows(rows *sql.Rows) ([]Movie, error) {
	var movies []Movie
	for rows.Next() {
		var movie Movie
		err := rows.Scan(&movie.Id, &movie.Name, &movie.Date, &movie.Tagline, &movie.Description, &movie.Minute, &movie.Rating)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	rows.Close()
	return movies, nil
}

func ToJsonMovie(movies []Movie) ([]byte, error) {
	result, err := json.Marshal(movies)
	if err != nil {
		return nil, err
	}
	return result, nil
}
