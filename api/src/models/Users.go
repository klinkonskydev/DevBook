package models

import (
	"errors"
	"strings"
	"time"
)

// User is a user that uses the platform
type User struct {
	ID           uint32    `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Nickname     string    `json:"nickname,omitempty"`
	Email        string    `json:"email,omitempty"`
	PasswordHash string    `json:"passwordHash,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
}

// Prepare will validate and formate the fields
func(u *User) Prepare() error {
	if err := u.validate(); err != nil {
		return err
	}

	u.formate()
	return nil
}

func(u *User) validate() error {
	if u.Name == "" {
		return errors.New("name is a required field")
	}

	if u.Email == "" {
		return errors.New("email is a required field")
	}

	if u.Nickname == "" {
		return errors.New("nickanme is a required field")
	}

	if u.PasswordHash == "" {
		return errors.New("password is a required field")
	}

	return nil
}

func (u *User) formate() {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.Nickname = strings.TrimSpace(u.Nickname)
}
