package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//e.Static("/assets", "public/assets")

	e.File("/", "public/index.html")

	api := e.Group("/api")
	api.GET("/hello", hello)

	return e
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
