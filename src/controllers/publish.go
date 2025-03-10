package controllers

import (
	"backend/src/answer"
	"backend/src/auth"
	"backend/src/database"
	"backend/src/models"
	"backend/src/repo"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Create a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, erro := auth.GetUserIDFromToken(r)
	if erro != nil {
		answer.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		answer.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var post models.Post

	if erro = json.Unmarshal(bodyRequest, &post); erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
		return
	}

	post.AuthorID = userID

	if erro := post.Prepare(); erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	defer db.Close()

	repo := repo.NewRepoPosts(db)

	post.ID, erro = repo.CreatePost(post)
	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return

	}

	answer.JSON(w, http.StatusCreated, post)

}

// Get a list of post
func GetPost(w http.ResponseWriter, r *http.Request) {

	userID, erro := auth.GetUserIDFromToken(r)
	if erro != nil {
		answer.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repo.NewRepoPosts(db)
	posts, erro := repo.GetPosts(userID)
	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return

	}

	answer.JSON(w, http.StatusOK, posts)

}

// Get a single post
func GetPostByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, erro := strconv.ParseUint(params["postid"], 10, 64)
	if erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repo.NewRepoPosts(db)
	post, erro := repo.GetPost(postID)
	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return

	}

	answer.JSON(w, http.StatusOK, post)

}

// Update the data of a post
func UpdatePost(w http.ResponseWriter, r *http.Request) {

	userID, erro := auth.GetUserIDFromToken(r)
	if erro != nil {
		answer.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)
	postID, erro := strconv.ParseUint(params["postid"], 10, 64)
	if erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repo.NewRepoPosts(db)
	postFromDB, erro := repo.GetPost(postID)
	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return

	}

	if postFromDB.AuthorID != userID {
		answer.Erro(w, http.StatusForbidden, errors.New("user have not permission to update post"))
		return
	}

	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		answer.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var post models.Post

	if erro = json.Unmarshal(bodyRequest, &post); erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = post.Prepare(); erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repo.UpdatePost(postID, post); erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	answer.JSON(w, http.StatusOK, nil)

}

// Delete a post
func DeletePost(w http.ResponseWriter, r *http.Request) {

	userID, erro := auth.GetUserIDFromToken(r)
	if erro != nil {
		answer.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)
	postID, erro := strconv.ParseUint(params["postid"], 10, 64)
	if erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repo.NewRepoPosts(db)
	postFromDB, erro := repo.GetPost(postID)
	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return

	}

	if postFromDB.AuthorID != userID {
		answer.Erro(w, http.StatusForbidden, errors.New("user have not permission to delete post"))
		return
	}

	if erro = repo.DeletePostByID(postID); erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	answer.JSON(w, http.StatusOK, nil)
}
