package repository

import (
	"api/src/models"
	"database/sql"
)

type postsRepository struct {
	db *sql.DB
}

// PostsRepository is a repository for posts
func PostsRepository(db *sql.DB) *postsRepository {
	return &postsRepository{db}
}

// Create insert a new post in the database
func(postsRepository postsRepository) Create(post models.Post) (uint32, error)  {
	statement, err := postsRepository.db.Prepare(
		"INSERT INTO posts(title, content, author_id) VALUES(?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	row, err := statement.Exec(&post.Title, &post.Content, &post.AuthorID); 
	if err != nil {
		return 0, err
	}

	postID, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(postID), nil
}

// GetByID get a post in the database
func(postsRepository postsRepository) GetByID(userID uint32) (models.Post, error)  {
	row, err := postsRepository.db.Query(`
		SELECT p.*, u.nickname FROM 
		posts p INNER JOIN users u
		ON u.id = p.author_id WHERE p.id = ?
	`,
		userID,
	)
	if err != nil {
		return models.Post{}, nil
	}
	defer row.Close()

	var post models.Post
	if row.Next() {
		if err := row.Scan(
			&post.ID, 
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Upvotes,
			&post.CreatedAt,
			&post.AuthorNickname,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

// Get get all user connections posts
func(postsRepository postsRepository) Get(userID uint32) ([]models.Post, error) {
	rows, err := postsRepository.db.Query(`
		SELECT DISTINCT p.*, u.nickname FROM posts p 
		INNER JOIN users u ON u.id = p.author_id 
		INNER JOIN followers f ON p.author_id = f.user_id 
		WHERE u.id = ? OR f.follower_id = ?
		ORDER BY 1 desc
	`, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.ID, 
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Upvotes,
			&post.CreatedAt,
			&post.AuthorNickname,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// Update get all user connections posts
func(postsRepository postsRepository) Update(postID uint32, post models.Post) error {
	statement, err := postsRepository.db.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(&post.Title, &post.Content, &postID); err != nil {
		return err
	}

	return nil
}

// Delete get all user connections posts
func(postsRepository postsRepository) Delete(postID uint32) error {
	statement, err := postsRepository.db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(&postID); err != nil {
		return err
	}

	return nil
}

// GetPostsByUser get all user posts
func(postsRepository postsRepository) GetPostsByUser(userID uint32) ([]models.Post, error) {
	rows, err := postsRepository.db.Query(`
		SELECT p.*, u.nickname FROM posts p
		INNER JOIN users u
		ON u.id = p.author_id
		WHERE p.author_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.ID, 
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Upvotes,
			&post.CreatedAt,
			&post.AuthorNickname,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// Upvote add one upvote to a post
func(postsRepository postsRepository) Upvote(postID uint32) error {
	statement, err := postsRepository.db.Prepare("UPDATE posts SET upvotes = upvotes + 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

// Downvote remove one upvote to a post 
func(postsRepository postsRepository) Downvote(postID uint32) error {
	statement, err := postsRepository.db.Prepare(`
		UPDATE posts SET upvotes = 
		CASE 
			WHEN upvotes > 0 THEN upvotes - 1
		ELSE 
			0 END 
		WHERE id = ?
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}
