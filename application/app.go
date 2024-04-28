package application

import (
	"net/http"

	"github.com/bvrple/application/routes"
)

type App struct { 
	Router http.Handler
}

func New() *App {
	app := &App{
		router: loadRoutes()
	}

	return app
}