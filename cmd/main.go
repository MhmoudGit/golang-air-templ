package main

import (
	"dashboard/src/components/button"
	"dashboard/src/view/home"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var Usersx []home.User = []home.User{
	{Id: 1, Name: "Mahmoud", State: "Good"},
	{Id: 1, Name: "Ahmed", State: "Good"},
}

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
	component := home.Hello(Usersx, buttons)
	return component.Render(c.Request().Context(), c.Response())
}

// State form handler
func stateHandler(c echo.Context) error {
	// Create an instance of the User struct
	user := new(home.User)

	// Bind form data to the User struct
	if err := c.Bind(user); err != nil {
		return err
	}

	updateUserState(Usersx, user.Name, user.State)
	buttons := []home.Buttons{
		{Placeholder: "Good", Theme: button.Good},
		{Placeholder: "Danger", Theme: button.Danger},
		{Placeholder: "Warning", Theme: button.Warning},
	}
	component := home.Card(user.Name, user.State, buttons)
	return component.Render(c.Request().Context(), c.Response())
}

func updateUserState(users []home.User, targetName string, newState string) ([]home.User, error) {
	for i, user := range users {
		if user.Name == targetName {
			// Found the target user, update the state
			users[i].State = newState
			return users, nil
		}
	}

	// User not found
	return nil, fmt.Errorf("user with name %s not found", targetName)
}
