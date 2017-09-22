package main

import (
	"fmt"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/render"
	"github.com/go-chi/cors"
	"github.com/pressly/chi"
	"log"
	"net/http"
)

type result struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Data    string `json:"data"`
}

func (s *result) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func main() {
	fmt.Println("tada!")
	r := chi.NewRouter()
	corsOption := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(corsOption.Handler)
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/dummy", createAgreementHandler())

	log.Println("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createAgreementHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, &result{Success: true})
	})
}
