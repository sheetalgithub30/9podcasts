package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
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
func deleteCategories(c echo.Context) (err error) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	us, err := db.Exec(`Delete from categories where id = $1 `, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	return c.JSON(http.StatusOK, us)
}
func updateCategories(c echo.Context) (err error) {
	ct := &Category{}

	if err = c.Bind(ct); err != nil {
		return
	}

	q := `UPDATE categories SET title = $1 where id =$2`
	_, err = db.Exec(q, ct.Title, ct.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	return c.JSON(http.StatusOK, ct)

}
