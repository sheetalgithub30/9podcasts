package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type TokenCredentials struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func GenerateToken(email string) (string, error) {
	// Create a new random session token
	uuid, _ := uuid.NewUUID()
	token := uuid.String()
	// Set the sessionID in the cache, along with the user whom it represents
	// The sessionID has an expiry time in seconds
	seconds := strconv.Itoa(60 * 10)
	_, err = cache.Do("SETEX", token, seconds, email)
	return token, err
}

func DeleteToken(token string) error {
	// Delete the older session token
	_, err = cache.Do("DEL", token)
	return err
}

func GenerateLink(email string) (string, error) {
	result := db.QueryRow("SELECT email FROM users WHERE email=$1", email)
	temp := ""
	tempPtr := &temp
	err = result.Scan(tempPtr)

	// If an entry with the email does not exist, send an "Unauthorized"(401) status
	if err == sql.ErrNoRows {
		log.Println(err)
		return "", err
	}

	// If the error is of any other type, send a 500 status
	if err != nil {
		log.Println(err)
		return "", err
	}

	//if email exists then send the link after generating token
	token, err := GenerateToken(email)
	if err != nil {
		log.Println("error generating token")
		return "", err
	}

	domain := "http://localhost"
	link := domain + "/resetpass?token=" + token
	return link, err
}

func ForgotPassword(c echo.Context) (err error) {
	enteredCreds := &TokenCredentials{}
	if err = c.Bind(enteredCreds); err != nil {
		return
	}
	email := enteredCreds.Email
	log.Println("email=", email)
	link, err := GenerateLink(email)
	if err != nil {
		log.Println(err)
		return
	}
	if link == "" {
		log.Println("Empty Link")
		c.String(http.StatusOK, "Empty Link")
		return
	}

	//To be implemented
	SendMail()

	return c.String(http.StatusOK, link)
}

func SendMail() {

}

func ResetPassword(c echo.Context) (err error) {
	enteredCreds := &TokenCredentials{}
	if err = c.Bind(enteredCreds); err != nil {
		return
	}
	token := enteredCreds.Token
	log.Println("token=", token)

	// We then get the email of the user from our cache, where we set the token
	email, err := cache.Do("GET", token)
	if err != nil {
		// If there is an error fetching from cache, return an internal server error status
		return c.NoContent(http.StatusInternalServerError)
	}
	if email == nil {
		// If the session token is not present in cache, return an unauthorized error
		return c.NoContent(http.StatusUnauthorized)
	}

	return c.String(http.StatusOK, fmt.Sprintf("email associated with token= %s", email))
}
