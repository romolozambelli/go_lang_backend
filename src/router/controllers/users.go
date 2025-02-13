package controllers

import (
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

	fmt.Println("Creating a new user")

	bodyRequest, erro := io.ReadAll(r.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := database.Connect()

	if erro != nil {
		log.Fatal(erro)
	}

	repository := repo.NewRepoUsers(db)
	repository.CreateNewUser(user)

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

	fmt.Println("Updating a user")

	bodyRequest, erro := io.ReadAll(r.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := db.Connect()

	if erro != nil {
		log.Fatal(erro)
	}

}

// Delete user from an id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User by ID!"))
}
