package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/person/create", create)
	e.GET("/person/get", get)
	e.PUT("/person/update", updated)
	e.DELETE("/person/delete", delete)

	e.Start(":8082")
}

func create(c echo.Context) error {
	return c.String(http.StatusOK, "created")
}

func updated(c echo.Context) error {
	return c.String(http.StatusOK, "updated")
}

func delete(c echo.Context) error {
	return c.String(http.StatusOK, "deleted")
}

func get(c echo.Context) error {
	return c.String(http.StatusOK, "get")
}
