package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type episode struct {
	Id                   int64  `json:"id"`
	Title                string `json:"title"`
	Description          string `json:"description"`
	Season               int64  `json:"season"`
	Episode              int64  `json:"episode"`
	Type_of_episode      string `json:"type_of_episode"`
	IsExplicit           bool   `json:"isExplicit"`
	Episode_cover_art    string `json:"episode_cover_art"`
	Episode_content_path string `json:"episode_content_path"`
	Published            bool   `json:"published"`
	Created_at           string `json:"created_at"`
	Updated_at           string `json:"updated_at"`
}

func createEpisodes(c echo.Context) (err error) {

	ep := &episode{}

	if err = c.Bind(ep); err != nil {
		return
	}

	q := `INSERT INTO episodes (title,description,season,episode,type_of_episode,isExplicit,episode_cover_art,episode_content_path,published,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING episode_id`
	err = db.QueryRow(q, ep.Title, ep.Description, ep.Season, ep.Episode, ep.Type_of_episode, ep.IsExplicit, ep.Episode_cover_art, ep.Episode_content_path, ep.Published, ep.Created_at, ep.Updated_at).Scan(&ep.Id)
	if err != nil {
		fmt.Println(err)
		return
	}

	return c.JSON(http.StatusOK, ep)
}

func getEpisode(c echo.Context) (err error) {

	var eps []episode

	q := `SELECT * FROM episodes`

	rows, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var ep episode

		err = rows.Scan(&ep.Id, &ep.Title, &ep.Description, &ep.Season, &ep.Episode, &ep.Type_of_episode, &ep.IsExplicit, &ep.Episode_cover_art, &ep.Episode_content_path, &ep.Published, &ep.Created_at, &ep.Updated_at)
		if err != nil {
			fmt.Println(err)
		}

		eps = append(eps, ep)
	}

	return c.JSON(http.StatusOK, eps)
}
