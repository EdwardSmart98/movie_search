package apiServer

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"log"
	"movieInfo/internal/database"
	"movieInfo/internal/routing"
	"net/http"
)

func Run() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", homeHandler)

	routing.AddRoutes(r)

	/// Make a test connection to the server
	connection := database.GetConnection()
	err := connection.Ping()
	if err != nil {
		println("Connection failed to server")
		println(err.Error())
		log.Panicf(err.Error())
	} else {
		println("Connection successful, ready to serve requests")
		connection.Close()
	}
	http.ListenAndServe(":8080", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	ctx := make(map[string]string)
	ctx["Name"] = "Neeraj"
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		println("Error parsing template")
		println(err.Error())
		return
	}
	print("template text", t.Name())
	err = t.ExecuteTemplate(w, "index.html", ctx)
	if err != nil {
		print("Error executing template")
		print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
