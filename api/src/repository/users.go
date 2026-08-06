// Package repository represents the server-to-database layer of the application
package repository

import (
	"database/sql"
	"fmt"

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
func (repo users) CreateUser(user models.User) (uint32, error) {
	statement, err := repo.db.Prepare(
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

// GetUsers get all users with includes name or nickname from database
func (repo users) GetUsers(nameOrNickname string) ([]models.User, error) {
	// Don't use %like% in SQL, instead of like you must use Full-Text Search to improve the performance
	// See the sql folder and create the Full-Text Search if it not created.
	rows, err := repo.db.Query(
		"SELECT id, name, nickname, email, createdAt FROM users WHERE MATCH(name, nickname) AGAINST(?)",
		nameOrNickname,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User
		if err = rows.Scan(
			&u.ID,
			&u.Name,
			&u.Nickname,
			&u.Email,
			&u.CreatedAt,
		); err != nil {
			fmt.Println("error 2")
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}
