package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nlgtEA/go-account/db/database"
	"github.com/nlgtEA/go-account/handler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Cannot load env. Error %v", err.Error())
	}

	dbURL := os.Getenv("DB_URL")
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}
	queries := database.New(conn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	userHandler := handler.UserHandler{}
	r.Get("/users", userHandler.HandlerUserShow)

	dashboardHander := handler.DashboardHandler{Queries: queries}
	r.Get("/", dashboardHander.HandleShow)

	fs := http.FileServer(http.Dir("view/public/"))
	r.Handle("/public/*", http.StripPrefix("/public/", fs))

	http.ListenAndServe(":3000", r)
}
