package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	env, ok := os.LookupEnv("ENV")
	log.Printf("ENV: %s", env)
	if !ok {
		log.Fatal("Environment variable ENV is required")
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"ENV": env,
		})
	})
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
