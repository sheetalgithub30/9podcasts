package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"

	"strings"
)

type Keyword struct {
	Item string `json:"item"`
}

func createKeywords(k echo.Context) (err error) {

	kw := &Keyword{}

	if err = k.Bind(kw); err != nil {
		return
	}

	if kw.Item == "" {
		return k.String(http.StatusBadRequest, "Item cannot be empty")
	}
	q := `INSERT INTO keywords (item) VALUES ($1) ON CONFLICT (item) DO NOTHING `
	_, err = db.Exec(q, strings.ToLower(kw.Item))
	if err != nil {
		fmt.Println(err)
		return
	}

	return k.JSON(http.StatusOK, kw)
}

func getKeywords(c echo.Context) (err error) {

	var kws []Keyword

	q := `SELECT item FROM keywords`

	rows, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var kw Keyword

		err = rows.Scan(&kw.Item)
		if err != nil {
			fmt.Println(err)
		}

		kws = append(kws, kw)
	}
	return c.JSON(http.StatusOK, kws)
}
