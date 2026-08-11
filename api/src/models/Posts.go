package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID             uint32    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorID       uint32    `json:"authorId,omitempty"`
	AuthorNickname string    `json:"authorNickname,omitempty"`
	Upvotes        uint64    `json:"upvotes"`
	CreatedAt      time.Time `json:"createdAt"`
}

// Prepare will validate and format
func (post *Post) Prepare(step string) error {
	if err := post.validate(); err != nil {
		return err
	}

	post.format(step)
	return nil
}

func (post *Post) validate() error {
	if strings.TrimSpace(post.Title) == "" {
		return errors.New("the title is required and can't be empty")
	}
	if strings.TrimSpace(post.Content) == "" {
		return errors.New("the content is required and can't be empty")
	}
	return nil
}

func (post *Post) format(step string) {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
	if step == "create" {
		post.CreatedAt = time.Now()
	}
}
