package main

import (
	"dashboard/view/home"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Static("/static", "static")

	app.GET("/", homeHandler)
	app.Start("localhost:3000")
}

func homeHandler(c echo.Context) error {
	component := home.Hello("Mahmoud")
	return component.Render(c.Request().Context(), c.Response())
}
