// Package repository represents the server-to-database layer of the application
package repository

import (
	"database/sql"

	"api/src/models"
)

type userRepository struct {
	db *sql.DB
}

// UsersRepository create an user repository
func UsersRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

// CreateUser insert a new user from database
func (userRepository userRepository) CreateUser(user models.User) (uint32, error) {
	statement, err := userRepository.db.Prepare(
		"INSERT INTO users (name, nickname, email, password) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(&user.Name, &user.Nickname, &user.Email, &user.Password)
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
func (userRepository userRepository) GetUsers(nameOrNickname string) ([]models.User, error) {
	// Don't use %like% in SQL, instead of like you must use Full-Text Search to improve the performance
	// See the sql folder and create the Full-Text Search if it not created.
	rows, err := userRepository.db.Query(
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
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

// GetUserByID get an user from the database where the userID is equals the database user id
func (userRepository userRepository) GetUserByID(userID uint32) (models.User, error) {
	rows, err := userRepository.db.Query(
		"SELECT id, name, nickname, email, createdAt FROM users WHERE id = ?",
		userID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var u models.User

	if rows.Next() {
		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Nickname,
			&u.Email,
			&u.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return u, nil
} 

// GetUserByID get an user from the database where the user.email is equals the database user email
func (userRepository userRepository) GetUserByEmail(email *string, password *string) (models.User, error) {
	row, err := userRepository.db.Query(
		"SELECT id, password FROM users WHERE email = ?",
		email,
	)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, nil
		}
	}

	return user, nil
}

// EditUser edit an user from the database
func (userRepository userRepository) EditUser(userID uint32, user models.User) error {
	statement, err := userRepository.db.Prepare(
		"UPDATE users SET name = ?, nickname = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nickname, &user.Email, userID); err != nil {
		return err
	}

	return nil
}

// DeleteUser remove an user from the database
func(userRepository userRepository) DeleteUser(userID uint32) error {
	statement, err := userRepository.db.Prepare(
		"DELETE FROM users WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID); err != nil {
		return err
	}
	
	return nil
}

// FollowUser will use a junction table for follower and user in the database
func(userRepository userRepository) FollowUser(userID uint32, followerID uint32) error {
	statement, err := userRepository.db.Prepare(
		"INSERT IGNORE INTO followers(user_id, follower_id) VALUES (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(&userID, &followerID); err != nil {
		return err
	}
	
	return nil
}

// UnfollowUser will remove the junction table between follower and user in the database
func(userRepository userRepository) UnfollowUser(userID uint32, followerID uint32) error {
	statement, err := userRepository.db.Prepare(
		"DELETE FROM followers WHERE user_id = ? AND follower_id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(&userID, &followerID); err != nil {
		return err
	}
	
	return nil
}

// Followers will return the followers from an user
func(userRepository userRepository) Followers(userID uint32) ([]models.User, error) {
	rows, err := userRepository.db.Query(`
		SELECT u.id, u.name, u.nickname, u.email, u.createdAt FROM users u 
		INNER JOIN followers f ON u.id = f.follower_id WHERE f.user_id = ?;
	`, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		if err := rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Following will return all users that a user follows
func(userRepository userRepository) Following(userID uint32) ([]models.User, error) {
	rows, err := userRepository.db.Query(`
		SELECT u.id, u.name, u.nickname, u.email, u.createdAt FROM users u
		INNER JOIN followers f ON u.id = f.user_id WHERE f.user_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err := rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
