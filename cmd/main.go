package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nlgtEA/go-account/handler"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	userHandler := handler.UserHandler{}
	r.Get("/users", userHandler.HandlerUserShow)

	fs := http.FileServer(http.Dir("view/public/"))
	r.Handle("/public/*", http.StripPrefix("/public/", fs))

	http.ListenAndServe(":3000", r)
}
