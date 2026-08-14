package models

import "time"

// Post represents a post made by an user
type Post struct {
	ID             uint32    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorID       uint32    `json:"authorId,omitempty"`
	AuthorNickname string    `json:"authorNickname,omitempty"`
	Upvotes        uint64    `json:"upvotes"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
}
