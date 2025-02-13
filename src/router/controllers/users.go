package controllers

import "net/http"

// Insert user on DB
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating User !"))
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
	w.Write([]byte("Updating User !"))
}

// Delete user from an id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User by ID!"))
}
