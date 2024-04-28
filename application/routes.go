package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	router.Route("/orders", loadOrderRoutes)

	return router
}

func loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{}

	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetById)
	router.Put("/{id}", orderHandler.UpdateById)
	router.Delete("/{id}", orderHandler.DeleteById)
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.Router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to listen to server: %w", err)
	}

	return nil
}
