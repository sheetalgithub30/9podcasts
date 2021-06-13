package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type episode struct {
	ID          int64  `json:"id"`
	PodcastID   int64  `json:"podcast_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	SeasonNo    int    `json:"season_no"`
	EpisodeNo   int    `json:"episode_no"`
	// 1 = Full | 2 = Bonus | 3 = trailer
	TypeOfEpisode    int       `json:"type_of_episode"`
	IsExplicit       bool      `json:"isExplicit"`
	EpisodeArtID     int64     `json:"episode_art_id"`
	EpisodeContentID int64     `json:"episode_content_id"`
	Published        bool      `json:"published"`
	PublishedAt      time.Time `json:"published_at"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func createEpisodes(c echo.Context) (err error) {

	ep := &episode{}

	if err = c.Bind(ep); err != nil {
		return
	}

	ep.CreatedAt = time.Now()
	ep.UpdatedAt = time.Now()

	q := `
		INSERT INTO episodes (
			title, description, season_no, episode_no, type_of_episode, 
			is_explicit, episode_art_id, episode_content_id, 
			published, created_at, updated_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id
	`
	err = db.QueryRow(q, ep.Title, ep.Description, ep.SeasonNo, ep.EpisodeNo, ep.TypeOfEpisode, ep.IsExplicit, ep.EpisodeArtID, ep.EpisodeContentID, ep.Published, ep.CreatedAt, ep.UpdatedAt).Scan(&ep.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	return c.JSON(http.StatusOK, ep)
}

func getEpisode(c echo.Context) (err error) {

	var eps []episode

	q := `SELECT id, title, description, season_no, episode_no,
				type_of_episode, is_explicit, episode_art_id, episode_content_id,
				published, created_at, updated_at FROM episodes
			`

	rows, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var ep episode

		err = rows.Scan(&ep.ID, &ep.Title, &ep.Description, &ep.SeasonNo, &ep.EpisodeNo, &ep.TypeOfEpisode, &ep.IsExplicit, &ep.EpisodeArtID, &ep.EpisodeContentID, &ep.Published, &ep.CreatedAt, &ep.UpdatedAt)
		if err != nil {
			fmt.Println(err)
		}

		eps = append(eps, ep)
	}

	return c.JSON(http.StatusOK, eps)
}
