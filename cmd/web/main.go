package main

import (
	"fmt"
	"go-web-server/pkg/config"
	"go-web-server/pkg/handlers"
	"go-web-server/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, error := render.CreateTemplateCache()
	if error != nil {
		log.Fatal("Cannot create template cache", error)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	//http.HandleFunc("/", handlers.Repo.HomeWebHandler)
	//http.HandleFunc("/about", handlers.Repo.AboutWebHandler)
	fmt.Println("Listening to port", portNumber)
	// _ allows us to ignore error from listener
	//_ = http.ListenAndServe(portNumber, nil)
	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err := serve.ListenAndServe()
	log.Fatal(err)
}
