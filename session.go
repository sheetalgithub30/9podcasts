package main

import (
	"strconv"

	"github.com/google/uuid"
)

func SetSessionID(email string) (string, error) {
	// Create a new random session token
	uuid, _ := uuid.NewUUID()
	sessionID := uuid.String()
	// Set the sessionID in the cache, along with the user whom it represents
	// The sessionID has an expiry time in seconds
	seconds := strconv.Itoa(60 * 2)
	_, err = cache.Do("SETEX", sessionID, seconds, email)
	return sessionID, err
}

func DeleteSessionID(sessionID string) error {
	// Delete the older session token
	_, err = cache.Do("DEL", sessionID)
	return err
}
