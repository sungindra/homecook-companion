package main

import (
	"homecook-companion/controller"

	"github.com/go-chi/chi/v5"
)

func handleRouting() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/manage", func(r chi.Router) {
		r.Route("/dish", func(r chi.Router) {
			r.Get("/", controller.ListDish)
			r.Get("/{id}", controller.GetDish)
			r.Get("/new", controller.NewDish)
			r.Post("/create", controller.CreateDish)
			r.Patch("/{id}", controller.UpdateDish)
			r.Delete("/{id}", controller.DeleteDish)
		})
	})
	return r
}
