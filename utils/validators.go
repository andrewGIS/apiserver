package utils

import (
	//"encoding/json"
	"errors"
	"net/http"
	"regexp"
)

// IsEmailValid checks if the email provided passes
// the required structure and length.
func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(
		"^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}" +
			"[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

// Function for general check of json
// Checking body type (only "application/json" is accepted)
func IsJsonValid(w http.ResponseWriter, r *http.Request) error {

	if r.Header.Get("Content-Type") != "" {
		value := r.Header.Get("Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return errors.New(msg)
		}
	}
	return nil
}
