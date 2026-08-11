package models

// Password represents the request payload from /update-password route
type Password struct {
	New string `json:"new"`
	Old string `json:"old"`
}
