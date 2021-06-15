package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

func (pd *podcast) GenerateRSS() (err error) {
	buf := &bytes.Buffer{}
	t := template.New("feed.rss")
	tp, err := t.ParseFiles("feed.rss")
	if err != nil {
		return
	}
	err = tp.Execute(buf, pd)
	if err != nil {
		return
	}

	err = writeToFile(pd.WebsiteAddress+".rss", buf.Bytes())
	return
}

func writeToFile(fileName string, data []byte) (err error) {
	f, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		return
	}
	fmt.Println("RSS feed created successfully.")
	return
}
