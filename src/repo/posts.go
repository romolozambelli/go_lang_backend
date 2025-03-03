package repo

import (
	"backend/src/models"
	"database/sql"
)

// Represent a repo of posts
type Posts struct {
	db *sql.DB
}

// Create a new db instance to post
func NewRepoPosts(db *sql.DB) *Posts {
	return &Posts{db}
}

// Insert a new post on the database
func (repoPost Posts) CreatePost(post models.Post) (uint64, error) {

	statement, erro := repoPost.db.Prepare(
		"INSERT INTO posts (title, text, autor_id) VALUES (?,?,?)",
	)

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(post.Title, post.Text, post.AuthorID)
	if erro != nil {
		return 0, erro
	}

	lastInsertedID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastInsertedID), nil

}
