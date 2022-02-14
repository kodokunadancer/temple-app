package main

import (
	"log"
	"temple-app/backend/pkg/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(logMiddleware)

	e.File("/", "../web/index.html")

	api := e.Group("/api")
	api.POST("/signup", handler.SignUp)

	return e
}

func logMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("before action")
		if err := next(c); err != nil {
			c.Error(err)
		}
		log.Println("after action")
		return nil
	}
}
