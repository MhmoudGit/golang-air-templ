package main

import (
	"dashboard/src/components/button"
	"dashboard/src/view/home"
	"dashboard/src/view/layout"
	usersView "dashboard/src/view/users"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var Usersx []usersView.User = []usersView.User{
	{Name: "Mahmoud", State: "Good"},
	{Name: "Ahmed", State: "Good"},
}

var buttons []usersView.Buttons = []usersView.Buttons{
	{Placeholder: "Good", Theme: button.Good},
	{Placeholder: "Danger", Theme: button.Danger},
	{Placeholder: "Warning", Theme: button.Warning},
}

var logged string = "off"

func main() {
	// Creating new instance of echo
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serving static files (css, html, images)
	e.Static("/static", "static")

	// Routes
	e.GET("/", mainHandler)
	e.GET("/home", homeHandler)
	e.GET("/login", loginHandler)
	e.GET("/logout", logoutHandler)
	e.GET("/users", usersHandler)
	e.GET("/tables", usersHandler)
	e.POST("/state", stateHandler)
	e.POST("/add", addUserHandler)

	// Starting echo server
	e.Start("localhost:3000")
}

func loginHandler(c echo.Context) error {
	logged = "on"
	return c.Redirect(302, "/")
}

func logoutHandler(c echo.Context) error {
	logged = "off"
	return c.Redirect(302, "/")
}

// Main Page Handler
func mainHandler(c echo.Context) error {
	if logged == "on" {
		return c.Redirect(302, "/home")
	}
	component := layout.Base("/", logged)
	return component.Render(c.Request().Context(), c.Response())
}

// Main Page Handler
func homeHandler(c echo.Context) error {
	component := home.Home(logged)
	return component.Render(c.Request().Context(), c.Response())
}

// Users Page Handler
func usersHandler(c echo.Context) error {

	component := usersView.UsersView(Usersx, buttons, logged)
	return component.Render(c.Request().Context(), c.Response())
}

// State form handler
func stateHandler(c echo.Context) error {
	// Create an instance of the User struct
	user := new(usersView.User)

	// Bind form data to the User struct
	if err := c.Bind(user); err != nil {
		return err
	}

	updateUserState(Usersx, user.Name, user.State)

	component := usersView.Card(user.Name, user.State, buttons)
	return component.Render(c.Request().Context(), c.Response())
}

// State form handler
func addUserHandler(c echo.Context) error {
	// Create an instance of the User struct
	var user usersView.User

	// Bind form data to the User struct
	if err := c.Bind(&user); err != nil {
		return err
	}

	Usersx = append(Usersx, user)

	component := usersView.UsersData(Usersx, buttons)
	return component.Render(c.Request().Context(), c.Response())
}

func updateUserState(users []usersView.User, targetName string, newState string) ([]usersView.User, error) {
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
