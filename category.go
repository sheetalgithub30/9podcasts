package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Category struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

func createCategory(c echo.Context) (err error) {

	ct := &Category{}

	if err = c.Bind(ct); err != nil {
		return
	}

	q := `INSERT INTO categories (title) VALUES ($1) RETURNING id`
	err = db.QueryRow(q, ct.Title).Scan(&ct.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	return c.JSON(http.StatusOK, ct)
}

func getCategories(c echo.Context) (err error) {

	var cts []Category

	q := `SELECT id, title FROM categories`

	rows, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var ct Category

		err = rows.Scan(&ct.ID, &ct.Title)
		if err != nil {
			fmt.Println(err)
		}

		cts = append(cts, ct)
	}

	return c.JSON(http.StatusOK, cts)
}
