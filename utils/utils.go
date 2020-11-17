package utils

import (
	"encoding/json"
	"net/http"
)

// Utils defines the shape of utils
type Utils struct {}

// NewUtils returns the utils struct
func NewUtils() *Utils {
	return &Utils{}
}

// ParseForm parses the req body
func (u *Utils) ParseForm(r *http.Request, dst interface{}) error {
	return json.NewDecoder(r.Body).Decode(dst)
}
