package main

import (
	"fmt"
	"log"
	"time"

	"net/http"

	"github.com/labstack/echo"
)

func Dashboard(c echo.Context) error {
	ck, err := c.Cookie("SessionID")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			// log.Println("error getting cookie")
			return c.NoContent(http.StatusUnauthorized)
		}
		// For any other type of error, return a bad request status
		return c.NoContent(http.StatusBadRequest)
	}
	sessionID := ck.Value

	// We then get the email of the user from our cache, where we set the sessionID
	res, err := cache.Do("GET", sessionID)
	if err != nil {
		// If there is an error fetching from cache, return an internal server error status
		return c.NoContent(http.StatusInternalServerError)
	}
	if res == nil {
		// If the session token is not present in cache, return an unauthorized error
		// log.Println("session token not in cache")
		return c.NoContent(http.StatusUnauthorized)
	}
	return c.String(http.StatusOK, fmt.Sprintf("Welcome %s!", res))
}

func refreshToken(c echo.Context) error {
	ck, err := c.Cookie("SessionID")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			// log.Println("error getting cookie")
			return c.NoContent(http.StatusUnauthorized)
		}
		// For any other type of error, return a bad request status
		return c.NoContent(http.StatusBadRequest)
	}
	sessionID := ck.Value

	// We then get the email of the user from our cache, where we set the sessionID
	res, err := cache.Do("GET", sessionID)
	if err != nil {
		// If there is an error fetching from cache, return an internal server error status
		return c.NoContent(http.StatusInternalServerError)
	}
	if res == nil {
		// If the session token is not present in cache, return an unauthorized error
		// log.Println("session token not in cache")
		return c.NoContent(http.StatusUnauthorized)
	}

	// Now, create a new session token for the current user
	NewSessionID, err := SetSessionID(fmt.Sprintf("%s", res))
	if err != nil {
		log.Println(err)
		return err
	}

	// Delete the older session token
	err = DeleteSessionID(sessionID)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	// Set the new token as the users `SessionID` cookie
	cookie := new(http.Cookie)
	cookie.Name = "SessionID"
	cookie.Value = NewSessionID
	cookie.Expires = time.Now().Add(60 * time.Second)
	c.SetCookie(cookie)

	return c.String(http.StatusOK, NewSessionID)
}
