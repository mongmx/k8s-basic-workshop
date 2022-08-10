package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := dbConnect()
	if err != nil {
		log.Fatal("DB connect error", err)
	}
	err = db.AutoMigrate(&Product{})
	if err != nil {
		log.Fatal("AutoMigrate error", err)
	}
	db.Create(&Product{Code: "D42", Price: 100})

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"name":    "My App API",
			"version": "x.x.x",
		})
	})
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	e.GET("/products", func(c echo.Context) error {
		var products []Product
		db.Find(&products)
		return c.JSON(http.StatusOK, products)
	})
	e.Logger.Fatal(e.Start(":8080"))
}

func dbConnect() (*gorm.DB, error) {
	pgHost, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		return nil, fmt.Errorf("Environment variable POSTGRES_HOST is required")
	}
	pgUser, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		return nil, fmt.Errorf("Environment variable POSTGRES_USER is required")
	}
	pgPassword, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("Environment variable POSTGRES_PASSWORD is required")
	}
	pgDBName, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		return nil, fmt.Errorf("Environment variable POSTGRES_DB is required")
	}

	dsn := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s port=%s sslmode=%s",
		pgHost,
		pgDBName,
		pgUser,
		pgPassword,
		"5432",
		"disable",
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
