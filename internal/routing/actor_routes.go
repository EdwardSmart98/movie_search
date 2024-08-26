package routing

import (
	"github.com/go-chi/chi/v5"
	"movieInfo/internal/models"
	"movieInfo/internal/repository"
	"net/http"
	"strconv"
)

func AddActorRoutes(r *chi.Mux) {
	r.Get("/actor", getActors)
	r.Get("/actor/{id}", getActor)
	r.Get("/actor/search", searchActors)
}

func getActors(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Actor"))
	if err != nil {
		return
	}
}

func getActor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		_, err := w.Write([]byte("Id must be an integer"))
		if err != nil {
			return
		}
		return
	}
	actorObject, err := repository.GetActorByID(idInt)
	if err != nil {
		_, err := w.Write([]byte("Error getting actor" + err.Error()))
		if err != nil {
			return
		}
		return
	}
	jsonActor, err := models.ToJsonActor([]models.Actor{actorObject})
	if err != nil {
		_, err := w.Write([]byte("Error converting actor to json" + err.Error()))
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

func searchActors(w http.ResponseWriter, r *http.Request) {

	searchString := r.URL.Query().Get("search")
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		_, err := w.Write([]byte("Page must be an integer"))
		if err != nil {
			return
		}
		return
	}
	actors, err := repository.SearchActors(searchString, pageInt)
	if err != nil {
		_, err := w.Write([]byte("Error searching for actors" + err.Error()))
		if err != nil {
			return
		}
		return
	}
	jsonActors, err := models.ToJsonActor(actors)
	if err != nil {
		_, err := w.Write([]byte("Error converting actors to json" + err.Error()))
		if err != nil {
			return
		}
	}
	_, err = w.Write(jsonActors)
	if err != nil {
		return
	}
	return
}
