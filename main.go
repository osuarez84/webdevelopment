package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/osuarez84/webdevelopment/controllers"
	"github.com/osuarez84/webdevelopment/views"
	"log"
	"net/http"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	t, err := views.Parse(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	tpl, err := views.Parse("templates/home.gohtml")
	if err != nil {
		panic(err)
	}

	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse("templates/contact.gohtml")
	if err != nil {
		panic(err)
	}

	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse("templates/faq.gohtml")
	if err != nil {
		panic(err)
	}

	r.Get("/faq", controllers.StaticHandler(tpl))
	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", r)
}
