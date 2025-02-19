package controllers

import (
	"backend/src/answer"
	"backend/src/database"
	"backend/src/models"
	"backend/src/repo"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Insert user on DB
func CreateUser(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("controllers/users --> Creating a new ...\n")

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

	fmt.Printf("controllers/users --> User -> %d Created with success \n", user.ID)
	answer.JSON(w, http.StatusCreated, user)
}

// Get all users from db
func GetUsers(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("controllers/users --> Get users by nick or name ...\n")

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

	fmt.Printf("controllers/users --> Get users by a given id ...\n")

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

	fmt.Printf("controllers/users --> Updating an existent user\n")

	param := mux.Vars(r)

	userID, erro := strconv.ParseUint(param["userID"], 10, 64)
	if erro != nil {
		answer.Erro(w, http.StatusBadRequest, erro)
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

	fmt.Printf("controllers/users --> User updated ID -> %d with success\n", userID)
	answer.JSON(w, http.StatusNoContent, nil)

}

// Delete user from an id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("controllers/users --> Deleting an existent user\n")

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
	if erro := repo.DeleteUserByID(userID); erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	fmt.Printf("controllers/users --> Deleted user ID -> %d with success\n", userID)
	answer.JSON(w, http.StatusNoContent, nil)

}
