package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo"
)

var (
	db  *sql.DB
	err error
)

func main() {
	connectionString := os.Getenv("CONN")
	if connectionString == "" {
		log.Fatal("error getting connection string")
	}

	db, err = sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	api := echo.New()

	api.GET("/", rootHandler)

	api.POST("/categories", createCategory)
	api.GET("/categories", getCategories)
	api.POST("/keywords", createKeywords)
	api.GET("/keywords", getKeywords)
	api.POST("/episodes", createEpisodes)
	api.GET("/episodes", getEpisode)
	api.POST("/podcasts", createPodcast)
	api.GET("/podcasts", getPodcast)

	api.HideBanner = true
	api.Start(":9999")
}

func rootHandler(c echo.Context) (err error) {

	return c.String(http.StatusOK, "Welcome to 9Podcasts Api")
}
