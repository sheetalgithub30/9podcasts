package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type podcast struct {
	ID             int64     `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	WebsiteAddress string    `json:"website_address"`
	CategoryID     int64     `json:"category_id,omitempty"`
	Category       string    `json:"category,omitempty"`
	Language       string    `json:"language"`
	IsExplicit     bool      `json:"isExplicit"`
	CoverArtID     int64     `json:"cover_art_id,omitempty"`
	CoverArt       Media     `json:"cover_art_path,omitempty"`
	AuthorName     string    `json:"author_name"`
	AuthorEmail    string    `json:"author_email"`
	Copyright      string    `json:"copyright"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Episodes       []episode `json:"episodes,omitempty"`
}

func createPodcast(c echo.Context) (err error) {

	pd := &podcast{}

	if err = c.Bind(pd); err != nil {
		return
	}

	pd.CreatedAt = time.Now()
	pd.UpdatedAt = time.Now()

	q := `
		INSERT INTO podcasts(
			title, description, website_address, category_id, language, 
			is_explicit, cover_art_id, author_name, author_email, copyright, 
			created_at, updated_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING id
		`
	err = db.QueryRow(q, pd.Title, pd.Description, pd.WebsiteAddress, pd.CategoryID,
		pd.Language, pd.IsExplicit, pd.CoverArtID, pd.AuthorName, pd.AuthorEmail, pd.Copyright,
		pd.CreatedAt, pd.UpdatedAt).Scan(&pd.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	return c.JSON(http.StatusOK, pd)
}

func getPodcast(c echo.Context) (err error) {

	var pds []podcast

	q := `
		SELECT id, title, description, website_address, category_id, language, is_explicit,
		 cover_art_id, author_name, author_email,copyright, created_at, updated_at FROM podcasts
	`

	rows, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var pd podcast

		err = rows.Scan(&pd.ID, &pd.Title, &pd.Description, &pd.WebsiteAddress, &pd.CategoryID,
			&pd.Language, &pd.IsExplicit, &pd.CoverArtID, &pd.AuthorName, &pd.AuthorEmail, &pd.Copyright,
			&pd.CreatedAt, &pd.UpdatedAt)
		if err != nil {
			fmt.Println(err)
		}

		pd.CoverArt, _ = getMediaByID(pd.CoverArtID)

		pds = append(pds, pd)
	}

	return c.JSON(http.StatusOK, pds)
}

func getPodcastByID(c echo.Context) (err error) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return
	}

	q := `
	SELECT id, title, description, website_address, category_id, language, is_explicit,
	 cover_art_id, author_name, author_email,copyright, created_at,updated_at FROM podcasts
	WHERE id = $1
`
	var pd podcast
	err = db.QueryRow(q, id).Scan(&pd.ID, &pd.Title, &pd.Description, &pd.WebsiteAddress,
		&pd.CategoryID, &pd.Language, &pd.IsExplicit, &pd.CoverArtID, &pd.AuthorName, &pd.AuthorEmail,
		&pd.Copyright, &pd.CreatedAt, &pd.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		return
	}

	pd.Episodes, err = _getPodcastEpisodes(pd.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	pd.CoverArt, _ = getMediaByID(pd.CoverArtID)
	return c.JSON(http.StatusOK, pd)
}
