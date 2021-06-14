package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
	"time"
)

func GenerateRSS(pd podcast) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	t := template.New("feed.rss")
	tp, err := t.ParseFiles("feed.rss")
	if err != nil {
		log.Fatal("Error executing template :", err)
	}
	err = tp.Execute(buf, pd)
	if err != nil {
		log.Fatal("Execute error:", err)
	}
	return buf, err
}

func UpdateRSS() {

}

func GetPodcastDetails() {

}

//sample testing
func main1() {
	var pd podcast
	pd.AuthorEmail = "hello@Gmail.com"
	pd.AuthorName = "pk"
	pd.Category = "art"
	pd.Description = "this is podcast"
	pd.Title = "hello"
	pd.IsExplicit = false
	pd.Language = "en"

	var ed episode
	ed.Title = "first episode"
	ed.Description = "my first epidsode"
	ed.PublishedAt = time.Now()
	ed.IsExplicit = false
	pd.Episodes = []episode{ed, ed, ed}
	buf, _ := GenerateRSS(pd)
	fmt.Println(buf.String())
}
