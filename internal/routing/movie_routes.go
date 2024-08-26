package routing

import (
	"github.com/go-chi/chi/v5"
	"movieInfo/internal/errorHandling"
	"movieInfo/internal/models"
	"movieInfo/internal/repository"
	"movieInfo/internal/validation"
	"net/http"
	"strconv"
)

func AddMovieRoutes(r *chi.Mux) {
	r.Get("/movie", getMovies)
	r.Get("/movie/{id}", getMovie)
	r.Get("/movie/search", searchMovies)
	r.Get("/movie/{id}/actors", getActorsInMovie)
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Movie"))
	if err != nil {
		return
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		_, err := w.Write([]byte("Id must be an integer"))
		if err != nil {
			return
		}
		return
	}
	movieObject, err := repository.GetMovieByID(idInt)
	if err != nil {
		_, err := w.Write([]byte("Error getting movie" + err.Error()))
		if err != nil {
			return
		}
		return
	}
	jsonMovie, err := models.ToJsonMovie([]models.Movie{movieObject})
	if err != nil {
		_, err := w.Write([]byte("Error converting movie to json" + err.Error()))
		if err != nil {
			return
		}
	}
	_, err = w.Write(jsonMovie)
	if err != nil {
		return
	}
	return
}

func searchMovies(w http.ResponseWriter, r *http.Request) {

	searchQuery, err := validation.ValidateSearchQuery(r)
	if err != nil {
		errorHandling.SendError(w, errorHandling.InvalidParametersError([]string{"search", "page"}))
		return
	}
	moviesObject, err := repository.SearchMovies(searchQuery.SearchString, searchQuery.Page)

	if err != nil {
		errorHandling.SendError(w, errorHandling.DescribedInternalServerError("Error searching for movies"))
		return
	}
	jsonMovie, err := models.ToJsonMovie(moviesObject)
	if err != nil {
		errorHandling.SendError(w, errorHandling.DescribedInternalServerError("Error converting movies to json"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonMovie)
	return
	// Convert the json object into the html template and write it to the response writer
	/*	t, err := template.ParseFiles("./templates/movies/movie_list_template.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			println("Error parsing template")
			println(err.Error())
			return
		}

		data := template_data_models.MovieListPageData{
			Movies: moviesObject,
		}

		err = t.Execute(w, data)

		//_, err = w.Write(jsonMovie)
		if err != nil {
			println("Error writing json")
			println(err.Error())
			return
		}
		return*/
}

func getActorsInMovie(w http.ResponseWriter, r *http.Request) {
	movieId := chi.URLParam(r, "id")
	movieIdInt, err := strconv.Atoi(movieId)
	if err != nil {
		_, err := w.Write([]byte("Id must be an integer"))
		if err != nil {
			return
		}
		return
	}
	actorsObject, err := repository.GetActorsInAMovie(movieIdInt)
	if err != nil {
		_, err := w.Write([]byte("Error getting actors in movie" + err.Error()))
		if err != nil {
			return
		}
		return
	}
	jsonActor, err := models.ToJsonActor(actorsObject)
	if err != nil {
		_, err := w.Write([]byte("Error converting actors to json" + err.Error()))
		if err != nil {
			return
		}
	}
	_, err = w.Write(jsonActor)
	if err != nil {
		return
	}
	return
}
