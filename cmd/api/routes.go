package main

import (
	"bouguette/cmd/api/handlers"
)

func (app *Application) routes(handler handlers.Handler) {
	app.server.GET("/", handler.HealthCheck)
}
