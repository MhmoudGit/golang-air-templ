package main

import (
	"dashboard/src/components/button"
	"dashboard/src/view/home"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var State string

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
	e.POST("/state", stateHandler)

	// Starting echo server
	e.Start("localhost:3000")
}

// Home Page Handler
func homeHandler(c echo.Context) error {
	buttons := []home.Buttons{
		{Placeholder: "Good", Theme: button.Good},
		{Placeholder: "Danger", Theme: button.Danger},
		{Placeholder: "Warning", Theme: button.Warning},
	}
	component := home.Hello("Mahmoud", State, buttons)
	return component.Render(c.Request().Context(), c.Response())
}

// State form handler
func stateHandler(c echo.Context) error {
	State = c.FormValue("state")
	buttons := []home.Buttons{
		{Placeholder: "Good", Theme: button.Good},
		{Placeholder: "Danger", Theme: button.Danger},
		{Placeholder: "Warning", Theme: button.Warning},
	}
	component := home.Card("Mahmoud", State, buttons)
	return component.Render(c.Request().Context(), c.Response())
}
