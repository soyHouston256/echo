package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	person := e.Group("/person")
	person.Use(middlewareLogPersons)
	person.POST("", create)
	person.GET("/:id", get)
	person.PUT("/:id", updated)
	person.DELETE("/:id", delete)

	e.Start(":8082")
}

func create(c echo.Context) error {
	return c.String(http.StatusOK, "created")
}

func updated(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "updated "+id)
}

func delete(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "deleted "+id)
}

func get(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "get "+id)
}

func middlewareLogPersons(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Middleware log persons")
		return next(c)
	}
}
