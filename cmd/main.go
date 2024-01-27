package main

import (
	"fmt"
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

	http.ListenAndServe(":3000", r)

	fmt.Println("Hello I am working")
}
