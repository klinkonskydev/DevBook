package models

import (
	"errors"
	"strings"
	"time"

	"api/src/security"

	"github.com/badoux/checkmail"
)

// User is a user that uses the platform
type User struct {
	ID        uint32    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare will validate and formate the fields
func (u *User) Prepare(step string) error {
	if err := u.validate(step); err != nil {
		return err
	}

	if err := u.formate(step); err != nil {
		return err
	}

	return nil
}

func (u *User) validate(step string) error {
	if u.Name == "" {
		return errors.New("name is a required field")
	}

	if u.Email == "" {
		return errors.New("email is a required field")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("email format is not valid")
	}

	if u.Nickname == "" {
		return errors.New("nickanme is a required field")
	}

	if step == "register" && u.Password == "" {
		return errors.New("password is a required field")
	}

	return nil
}

func (u *User) formate(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.Nickname = strings.TrimSpace(u.Nickname)
	if step == "register" {
		passwordHash, err := security.Hash(u.Password)
		if err != nil {
			return err
		}

		u.Password = string(passwordHash)
	}

	return nil
}
