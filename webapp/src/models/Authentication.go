package models

// Authentication contains the id and token of the authenticated user
type Authentication struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
