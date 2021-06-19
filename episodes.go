package main

import (
	"fmt"
	"net/http"
	"strconv"
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
	EpisodeArtID     int64     `json:"episode_art_id,omitempty"`
	EpisodeArt       Media     `json:"episode_art,omitempty"`
	EpisodeContentID int64     `json:"episode_content_id,omitempty"`
	EpisodeContent   Media     `json:"episode_content,omitempty"`
	Published        bool      `json:"published"`
	PublishedAt      time.Time `json:"published_at,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func createEpisodes(c echo.Context) (err error) {

	ep := &episode{}

	if err = c.Bind(ep); err != nil {
		return
	}

	if ep.PodcastID == 0 || ep.Title == "" || ep.Description == "" {
		return c.String(http.StatusBadRequest, "Insufficient fields")
	}

	ep.CreatedAt = time.Now()
	ep.UpdatedAt = time.Now()

	q := `
		INSERT INTO episodes (
			podcast_id, title, description, season_no, episode_no, type_of_episode, 
			is_explicit, episode_art_id, episode_content_id, 
			published, created_at, updated_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING id
	`
	err = db.QueryRow(q, ep.PodcastID, ep.Title, ep.Description, ep.SeasonNo, ep.EpisodeNo,
		ep.TypeOfEpisode, ep.IsExplicit, ep.EpisodeArtID, ep.EpisodeContentID, ep.Published,
		ep.CreatedAt, ep.UpdatedAt).Scan(&ep.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	pd, _ := _getPodcastByID(ep.PodcastID)
	err = pd.GenerateRSS()
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, ep)
}

func getEpisodes(c echo.Context) (err error) {

	podcastID, err := strconv.Atoi(c.QueryParam("podcast_id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "podcast_id is required")
	}

	eps, err := _getPodcastEpisodes(int64(podcastID))
	if err != nil {
		fmt.Println(err)
		return
	}

	return c.JSON(http.StatusOK, eps)
}

func _getPodcastEpisodes(podcastID int64) (eps []episode, err error) {
	q := `SELECT id, podcast_id, title, description, season_no, episode_no,
	type_of_episode, is_explicit, episode_art_id, episode_content_id,
	published, created_at, updated_at FROM episodes WHERE podcast_id = $1
`
	rows, err := db.Query(q, podcastID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ep episode

		err = rows.Scan(&ep.ID, &ep.PodcastID, &ep.Title, &ep.Description, &ep.SeasonNo,
			&ep.EpisodeNo, &ep.TypeOfEpisode, &ep.IsExplicit, &ep.EpisodeArtID, &ep.EpisodeContentID,
			&ep.Published, &ep.CreatedAt, &ep.UpdatedAt)

		if err != nil {
			fmt.Println(err)
			continue
		}
		ep.EpisodeArt, err = getMediaByID(ep.EpisodeArtID)

		if err != nil {
			fmt.Println(err)
			continue
		}

		ep.EpisodeContent, err = getMediaByID(ep.EpisodeContentID)

		if err != nil {
			fmt.Println(err)
			continue
		}
		eps = append(eps, ep)
	}
	return eps, nil
}

func deleteEpisode(c echo.Context) (err error) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	us, err := db.Exec(`Delete from episodes where id = $1 `, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	return c.JSON(http.StatusOK, us)
}

func updateEpisode(c echo.Context) (err error) {
	ep := &episode{}

	ep.UpdatedAt = time.Now()

	if err = c.Bind(ep); err != nil {
		return
	}

	q := `UPDATE episodes SET title = COALESCE(NULLIF($1,''),title) ,
	description = COALESCE(NULLIF($2,''),description), 
	season_no= COALESCE(NULLIF($3,0),season_no), 
	episode_no = COALESCE(NULLIF($4,0),episode_no),
	type_of_episode = COALESCE(NULLIF($5,0),type_of_episode),
	is_explicit = COALESCE(NULLIF($6,false),is_explicit),
	episode_art_id = COALESCE(NULLIF($7,0),episode_art_id),
	episode_content_id = COALESCE(NULLIF($8,0),episode_content_id),
	published = COALESCE(NULLIF($9,false),published), 
	updated_at=$10 
	WHERE id = $11`

	_, err = db.Exec(q, ep.Title, ep.Description, ep.SeasonNo, ep.EpisodeNo, ep.TypeOfEpisode,
		ep.IsExplicit, ep.EpisodeArtID, ep.EpisodeContentID, ep.Published, ep.UpdatedAt, ep.ID)

	if err != nil {
		fmt.Println(err)
		return
	}
	return c.JSON(http.StatusOK, ep)
}
