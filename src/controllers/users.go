package controllers

import (
	"backend/src/answer"
	"backend/src/auth"
	"backend/src/database"
	"backend/src/models"
	"backend/src/repo"
	"backend/src/security"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Insert user on DB
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := io.ReadAll(r.Body)

	if erro != nil {
		answer.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("register"); erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()

	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repo.NewRepoUsers(db)
	user.ID, erro = repository.CreateNewUser(user)

	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	answer.JSON(w, http.StatusCreated, user)
}

// Get all users from db
func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("users"))

	db, erro := database.Connect()

	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repo.NewRepoUsers(db)
	users, erro := repo.GetUserOrNick(nameOrNick)
	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	answer.JSON(w, http.StatusOK, users)
}

// Get a user for a given id on db
func GetUserById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userID, erro := strconv.ParseUint(param["userID"], 10, 64)

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

	repo := repo.NewRepoUsers(db)

	user, erro := repo.SearchUserByID(userID)

	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	answer.JSON(w, http.StatusOK, user)

}

// Update user by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userID, erro := strconv.ParseUint(param["userID"], 10, 64)
	if erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
		return
	}

	tokenUserID, erro := auth.GetUserIDFromToken(r)
	if erro != nil {
		answer.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if tokenUserID != userID {
		answer.Erro(w, http.StatusForbidden, errors.New("not allowed to update other users"))
		return
	}

	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		answer.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("update"); erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	defer db.Close()

	repo := repo.NewRepoUsers(db)
	if erro := repo.UpdateUser(userID, user); erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	answer.JSON(w, http.StatusNoContent, nil)

}

// Delete user from an id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userID, erro := strconv.ParseUint(param["userID"], 10, 64)
	if erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
		return
	}

	tokenUserID, erro := auth.GetUserIDFromToken(r)
	if erro != nil {
		answer.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if tokenUserID != userID {
		answer.Erro(w, http.StatusForbidden, errors.New("not allowed to delete other users"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	defer db.Close()

	repo := repo.NewRepoUsers(db)
	if erro := repo.DeleteUserByID(userID); erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	answer.JSON(w, http.StatusNoContent, nil)

}

// Follow a user permanentely
func FollowUser(w http.ResponseWriter, r *http.Request) {

	userID, erro := auth.GetUserIDFromToken(r)

	if erro != nil {
		answer.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	param := mux.Vars(r)

	followID, erro := strconv.ParseUint(param["userID"], 10, 64)
	if erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if followID == userID {
		answer.Erro(w, http.StatusForbidden, errors.New("not possible to follow same user"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	defer db.Close()

	repo := repo.NewRepoUsers(db)
	if erro = repo.Follow(userID, followID); erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	answer.JSON(w, http.StatusNoContent, nil)
}

// Stop to follow a user
func UnFollowUser(w http.ResponseWriter, r *http.Request) {

	userID, erro := auth.GetUserIDFromToken(r)

	if erro != nil {
		answer.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	param := mux.Vars(r)

	followerID, erro := strconv.ParseUint(param["userID"], 10, 64)
	if erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if followerID == userID {
		answer.Erro(w, http.StatusForbidden, errors.New("not possible to unfollow same user"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	defer db.Close()

	repo := repo.NewRepoUsers(db)
	if erro = repo.UnFollow(userID, followerID); erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	answer.JSON(w, http.StatusNoContent, nil)

}

func GetFollowers(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userID, erro := strconv.ParseUint(param["userID"], 10, 64)
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

	repo := repo.NewRepoUsers(db)
	followers, erro := repo.GetFollowers(userID)

	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return

	}

	answer.JSON(w, http.StatusOK, followers)

}

// Get all the users a user is following
func GetFollowingUsers(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userID, erro := strconv.ParseUint(param["userID"], 10, 64)
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

	repo := repo.NewRepoUsers(db)

	users, erro := repo.GetFollowingUsers(userID)

	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return

	}

	answer.JSON(w, http.StatusOK, users)

}

// Update the user password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {

	tokenUserID, erro := auth.GetUserIDFromToken(r)
	if erro != nil {
		answer.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	param := mux.Vars(r)

	userID, erro := strconv.ParseUint(param["userID"], 10, 64)
	if erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if tokenUserID != userID {
		answer.Erro(w, http.StatusForbidden, errors.New("not allowed to update other users"))
		return
	}

	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		answer.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var password models.Password

	if erro = json.Unmarshal(bodyRequest, &password); erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	defer db.Close()

	repo := repo.NewRepoUsers(db)

	passwordDB, erro := repo.GetPassword(userID)
	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.CheckPassword(passwordDB, password.Password); erro != nil {
		answer.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	passwordHash, erro := security.Hash(password.NewPassword)
	if erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repo.UpdatePassword(userID, string(passwordHash)); erro != nil {

	}

	answer.JSON(w, http.StatusNoContent, nil)
}
