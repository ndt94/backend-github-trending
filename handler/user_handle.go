package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleSignIn(c echo.Context) error  {
	return c.JSON(http.StatusOK, echo.Map{
		"user": "jessenguyen",
		"email": "ndt94hvnh@gmail.com",
	})
}

func HandleSignUp(c echo.Context) error  {
	// Can custom name of struct when return to JSON using `json: "custom_name"`
	type User struct {
		Email string
		FullName string
		Age int
	}

	user := User{
		Email: "ndt94hvnh@gmail.com",
		FullName: "Jesse Nguyen",
		Age: 26,
	}
	return c.JSON(http.StatusOK, user)
}
