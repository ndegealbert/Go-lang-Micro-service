package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"

)

func routes() http.Handler {

	r := chi.NewRouter()
	
	r.Use(r.Middlewares()...)

	return r;

}
