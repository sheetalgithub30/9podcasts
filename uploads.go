package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

const uploadDir = "uploads"

type Media struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

func uploadMedia(c echo.Context) (err error) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}

	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer src.Close()

	// aman.dsom.png
	splitName := strings.Split(file.Filename, ".")
	fileName := uuid.New().String() + "." + splitName[len(splitName)-1]

	// Destination file
	dst, err := os.Create(filepath.Join(uploadDir, fileName))
	if err != nil {
		fmt.Println(dst)
		return
	}

	// copy from src to dst
	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println(err)
		return err
	}
	dst.Close()

	m := Media{
		Name: file.Filename,
		Path: fileName,
	}
	q := `INSERT INTO uploads(name,path) VALUES($1, $2) RETURNING id`

	err = db.QueryRow(q, m.Name, m.Path).Scan(&m.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	return c.JSON(http.StatusOK, m)
}

func getMediaByID(id int64) (m Media, err error) {
	q := `SELECT id, name, path from uploads where id = $1`
	err = db.QueryRow(q, id).Scan(&m.ID, &m.Name, &m.Path)
	return
}

func getMediaFile(c echo.Context) (err error) {
	p := c.Request().URL.Path
	p = strings.TrimPrefix(p, "/media")

	return c.File(uploadDir + p)
}
