package main

import (
	"dashboard/view/home"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Creating new instance of echo
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serving static files (css, html, images)
	e.Static("/static", "static")

	// Routes
	e.GET("/", homeHandler)

	// Starting echo server
	e.Start("localhost:3000")
}

// Home Page Handler
func homeHandler(c echo.Context) error {
	component := home.Hello("Mahmoud")
	return component.Render(c.Request().Context(), c.Response())
}
