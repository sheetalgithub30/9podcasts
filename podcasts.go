package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type podcast struct {
	Id             int64  `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	WebsiteAddress string `json:"website_address"`
	Category_id    int64  `json:"category_id"`
	Language       string `json:"language"`
	IsExplicit     bool   `json:"isExplicit"`
	Cover_art_path string `json:"cover_art_path"`
	Author_name    string `json:"author_name"`
	Author_email   string `json:"author_email"`
	Copyright      string `json:"copyright"`
	Created_at     string `json:"created_at"`
	Updated_at     string `json:"updated_at"`
	Published_at   string `json:"published_at"`
}

func createPodcast(c echo.Context) (err error) {

	pd := &podcast{}

	if err = c.Bind(pd); err != nil {
		return
	}

	q := `INSERT INTO podcasts(title,description,website_address,category_id,languages,is_Explicit,cover_art_path,author_name,author_email,copyright,created_at,updated_at,published_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) RETURNING podcast_id`
	err = db.QueryRow(q, pd.Title, pd.Description, pd.WebsiteAddress, pd.Category_id, pd.Language, pd.IsExplicit, pd.Cover_art_path, pd.Author_name, pd.Author_email, pd.Copyright, pd.Created_at, pd.Updated_at, pd.Published_at).Scan(&pd.Id)
	if err != nil {
		fmt.Println(err)
		return
	}

	return c.JSON(http.StatusOK, pd)
}

func getPodcast(c echo.Context) (err error) {

	var pds []podcast

	q := `SELECT category_id,title,description,website_address,category_id,languages,is_Explicit,cover_art_path,author_name,author_email,copyright,created_at,updated_at,published_at FROM podcasts`

	rows, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var pd podcast

		err = rows.Scan(&pd.Id, &pd.Title, &pd.Description, &pd.WebsiteAddress, &pd.Category_id, &pd.Language, &pd.IsExplicit, &pd.Cover_art_path, &pd.Author_name, &pd.Author_email, &pd.Copyright, &pd.Created_at, &pd.Updated_at, &pd.Published_at)
		if err != nil {
			fmt.Println(err)
		}

		pds = append(pds, pd)
	}

	return c.JSON(http.StatusOK, pds)
}
