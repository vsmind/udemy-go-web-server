package main

import (
	"github.com/bmizerany/pat"
	"go-web-server/pkg/config"
	"go-web-server/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	// Create multiplexer
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.HomeWebHandler))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.AboutWebHandler))

	return mux
}
