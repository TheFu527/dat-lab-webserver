package main

import (
	"dat-lab-webserver/dat-lab-webserver/apigen"
	"dat-lab-webserver/dat-lab-webserver/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	var handler handler.Handler
	apigen.RegisterHandlers(e, &handler)

	// Start server
	e.Logger.Fatal(e.Start("0.0.0.0:80"))
}
