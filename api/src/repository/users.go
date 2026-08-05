// Package repository represents the server-to-database layer of the application
package repository

import (
	"database/sql"

	"api/src/models"
)

type users struct {
	db *sql.DB
}

// UsersRepository create an user repository
func UsersRepository(db *sql.DB) *users {
	return &users{db}
}

// CreateUser insert a new user from database
func (repository users) CreateUser(user models.User) (uint32, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO users (name, nickname, email, passwordHash) VALUES (?, ?, ?, ?)",
	)

	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(&user.Name, &user.Nickname, &user.Email, &user.PasswordHash)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(lastInsertID), nil
}
