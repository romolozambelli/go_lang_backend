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

// Get the posts by a given ID
func (repoPost Posts) GetPost(postID uint64) (models.Post, error) {

	line, erro := repoPost.db.Query(`
		SELECT p.*, u.nickname FROM posts p INNER JOIN users u
		ON u.id = p.autor_id WHERE p.id = ?`, postID)

	if erro != nil {
		return models.Post{}, erro
	}
	defer line.Close()

	var post models.Post

	if line.Next() {
		if erro = line.Scan(
			&post.ID,
			&post.Title,
			&post.Text,
			&post.AuthorID,
			&post.Likes,
			&post.CreateDate,
			&post.AuthorNick,
		); erro != nil {
			return models.Post{}, erro
		}
	}

	return post, nil
}

// Get all posts from a user
func (repoPost Posts) GetPosts(userID uint64) ([]models.Post, error) {

	lines, erro := repoPost.db.Query(`SELECT DISTINCT p.*, u.nickname FROM posts p 
	INNER JOIN users u ON u.id = p.autor_id
	INNER JOIN followers f ON p.autor_id = f.user_id
	WHERE u.id = ? OR f.follower_id = ?
	ORDER BY 1 DESC`,
		userID, userID,
	)

	if erro != nil {
		return nil, erro
	}

	var posts []models.Post

	for lines.Next() {

		var post models.Post

		if erro = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Text,
			&post.AuthorID,
			&post.Likes,
			&post.CreateDate,
			&post.AuthorNick,
		); erro != nil {
			return nil, erro
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// Update a post on the database
func (repoPost Posts) UpdatePost(postID uint64, post models.Post) error {

	statement, erro := repoPost.db.Prepare("UPDATE posts SET title = ?, text = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(post.Title, post.Text, postID); erro != nil {
		return erro
	}

	return nil
}

// Delete a Post from the user
func (repoPost Posts) DeletePostByID(postID uint64) error {

	statement, erro := repoPost.db.Prepare("DELETE FROM posts WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(postID); erro != nil {
		return erro
	}
	return nil
}

// Get all the posts from a given user
func (repoPost Posts) GetPostsByID(userID uint64) ([]models.Post, error) {

	lines, erro := repoPost.db.Query(`
	SELECT p.*, u.nickname FROM posts p 
	JOIN users u ON u.id = p.autor_id 
	WHERE p.autor_id = ?`, userID)

	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

	var posts []models.Post

	for lines.Next() {

		var post models.Post

		if erro = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Text,
			&post.AuthorID,
			&post.Likes,
			&post.CreateDate,
			&post.AuthorNick,
		); erro != nil {
			return nil, erro
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// Like a post from a user
func (repoPost Posts) LikePost(postID uint64) error {
	statement, erro := repoPost.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(postID); erro != nil {
		return erro
	}
	return nil
}

// Unlike a post from a user
func (repoPost Posts) UnlikePost(postID uint64) error {
	statement, erro := repoPost.db.Prepare(`
	UPDATE posts SET likes = 
	CASE WHEN likes > 0 THAN
	likes - 1 ELSE likes
	END
	WHERE id = ?
	`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(postID); erro != nil {
		return erro
	}
	return nil
}
