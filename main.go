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
	login := new(Login)
	if err := c.Bind(login); err != nil {
		return err
	}
	fmt.Println("login: %s", login.Name)
	checked := checkLogin(login.Name, login.Password, login.Onetime)
	fmt.Println("Checked: %bool", checked)
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

	fmt.Printf("toReturn : %s\n", res)
	return c.JSONPretty(http.StatusOK, res, " ")

}

func main() {
	// Hello world, the web server
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/api/login", login)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
