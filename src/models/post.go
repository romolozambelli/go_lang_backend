package models

import (
	"errors"
	"strings"
	"time"
)

// Post represent a post published by a user
type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Text       string    `json:"text,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreateDate time.Time `json:"created,omitempty"`
}

// Prepare the new post
func (post *Post) Prepare() error {
	if erro := post.validatePost(); erro != nil {
		return erro
	}

	post.format()
	return nil
}

func (post *Post) validatePost() error {

	if post.Title == "" {
		return errors.New("title should not be empty")
	}

	if post.Text == "" {
		return errors.New("text should not be empty")
	}

	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Text = strings.TrimSpace(post.Text)
}
