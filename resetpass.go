package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type TokenCredentials struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type EmailData struct {
	UserName        string
	BrowserName     string
	OperatingSystem string
	ContactSupport  string
	ActionUrl       string
	Year            string
}

func GenerateToken(id int64, email string) (string, error) {
	// Create a new random session token
	uuid, _ := uuid.NewUUID()
	token := "00" + strconv.FormatInt(id, 10) + "-" + uuid.String()
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
	result := db.QueryRow("SELECT id FROM users WHERE email=$1", email)
	var id int64
	idPtr := &id
	err = result.Scan(idPtr)

	// If an entry with the email does not exist, send an "Unauthorized"(401) status
	if err == sql.ErrNoRows {
		log.Println("email doesn't exists ", err)
		return "", err
	}

	// If the error is of any other type, send a 500 status
	if err != nil {
		log.Println(err)
		return "", err
	}

	//if email exists then send the link after generating token
	token, err := GenerateToken(id, email)
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

	// send email with generated link
	// err = SendEmail(email, link)

	if err != nil {
		log.Println("error sending email")
		return
	}
	log.Println("Email sent with link = " + link)
	return c.String(http.StatusOK, "link sent successfully ")
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

	idSplitStr := strings.Split(token, "-")
	idStr := idSplitStr[0]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("id=", id)

	// code to be written to redirect to change password webpage by and pass id as param
	// return c.Redirect(http.StatusMovedPermanently, "/change_password/:"+idStr)
	return c.String(http.StatusOK, fmt.Sprintf("%s", email))
}

func RenderTemplate(htmlFile string, ed EmailData) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	parsedTemplate, _ := template.ParseFiles(htmlFile)
	err := parsedTemplate.Execute(buf, ed)
	if err != nil {
		log.Println("Error executing template :", err)
	}
	//fmt.Println(buf)
	return buf, err
}

func GenerateEmail(link string, email string) (htmlStr string) {
	var ed EmailData
	ed.ActionUrl = link
	ed.BrowserName = "Chrome"
	ed.ContactSupport = "support@9podcast.com"
	ed.OperatingSystem = "Linux"

	result := db.QueryRow("SELECT name FROM users WHERE email=$1", email)
	var uname string
	unamePtr := &uname
	err = result.Scan(unamePtr)

	// If an entry with the email does not exist, send an "Unauthorized"(401) status
	if err == sql.ErrNoRows {
		log.Println("email doesn't exists ", err)
		return
	}

	// If the error is of any other type, send a 500 status
	if err != nil {
		log.Println(err)
		return
	}

	ed.UserName = uname

	currentTime := time.Now()
	ed.Year = strconv.Itoa(currentTime.Year())

	buf, err := RenderTemplate("./templates/reset_template.html", ed)
	if err != nil {
		log.Println("Error generating html file foe email", err)
		return
	}
	htmlStr = buf.String()
	return htmlStr
}

func SendEmail(toEmail, link string) (err error) {
	htmlStr := GenerateEmail(link, toEmail)
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSKEY")
	toList := []string{toEmail}
	host := "smtp.gmail.com"
	port := "587"

	// msg := "Link to reset your Password : " + link
	msg := htmlStr
	body := []byte(msg)
	auth := smtp.PlainAuth("", from, password, host)
	err = smtp.SendMail(host+":"+port, auth, from, toList, body)
	return err
}
