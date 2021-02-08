package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func login(c echo.Context) error {
	msg := ""
	isValid := false
	user := c.FormValue("username")
	pass := c.FormValue("password")

	fmt.Printf("Login Attempt: %s, %s\n", user, pass)
	checked := checkLogin(user, pass)

	if checked {
		isValid = true
		msg += "Yay!"
	} else {
		msg += "Invalid Username or Password. Please try Again."
	}

	res := &Content{
		Valid:   isValid,
		Message: msg,
	}

	return c.JSON(http.StatusOK, res)

}

func main() {
	// Hello world, the web server
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", login)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
