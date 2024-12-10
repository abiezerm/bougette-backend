package main

import (
	"bouguette/cmd/api/handlers"
	"bouguette/cmd/api/middlewares"
	"bouguette/common"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Application struct {
	logger  echo.Logger
	server  *echo.Echo
	handler handlers.Handler
}

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}

	db, err := common.NewSqliteConnection()
	if err != nil {
		e.Logger.Fatal(err)
	}

	h := handlers.Handler{DB: db}

	app := Application{
		logger:  e.Logger,
		server:  e,
		handler: h,
	}

	e.Use(middlewares.CustomMiddleware)
	e.Use(middleware.Logger())
	app.routes(h)

	port := os.Getenv("APP_PORT")
	appAddress := fmt.Sprintf("localhost:%s", port)
	e.Logger.Fatal(e.Start(appAddress))
}
