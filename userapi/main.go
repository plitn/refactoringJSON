package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"refactoring/handlers"
	"time"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]interface{}{"time": time.Now().String()})
	})

	r.Route("/api/v1/users", func(r chi.Router) {

		r.Get("/", handlers.SearchUsers)
		r.Post("/", handlers.CreateUser)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetUser)
			r.Patch("/", handlers.UpdateUser)
			r.Delete("/", handlers.DeleteUser)
		})
	})

	err := http.ListenAndServe(":3333", r)
	if err != nil {
		log.Println("listen and server error")
		return
	}
}
