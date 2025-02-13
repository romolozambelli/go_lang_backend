package controllers

import (
	"backend/src/answer"
	"backend/src/database"
	"backend/src/repo"
	"backend/src/router/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	w.Write([]byte("Getting all Users !"))
}

// Get a user for a given id on db
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user by id!"))
}

// Update user by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("controllers/users --> Updating a user\n")

	bodyRequest, erro := io.ReadAll(r.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}

	//db, erro := db.Connect()

	if erro != nil {
		log.Fatal(erro)
	}

}

// Delete user from an id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User by ID!"))
}
