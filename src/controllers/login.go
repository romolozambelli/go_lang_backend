package controllers

import (
	"backend/src/answer"
	"backend/src/auth"
	"backend/src/database"
	"backend/src/models"
	"backend/src/repo"
	"backend/src/security"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Responsible for login the user
func Login(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("controler/login --> Starting the Login ...\n")

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

	repo := repo.NewRepoUsers(db)
	userDataDB, erro := repo.SearchByEmail(user.Email)
	if erro != nil {
		answer.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.CheckPassword(userDataDB.Password, user.Password); erro != nil {
		answer.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, _ := auth.GenerateToken(userDataDB.ID)

	fmt.Printf("controler/login --> User fetch with success ...\n")

	fmt.Println(token)

	w.Write([]byte(token))
}
