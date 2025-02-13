package repo

import (
	"backend/src/router/models"
	"database/sql"
)

// Struct of the USER Repo
type Users struct {
	db *sql.DB
}

// Create a repo for the users
func NewRepoUsers(db *sql.DB) *Users {
	return &Users{db}
}

// Create a new user based on the user model
func (repoUser Users) CreateNewUser(user models.User) (uint64, error) {
	statement, erro := repoUser.db.Prepare(
		"INSERT INTO USERS (name,nickname,email,password) values (?,?,?,?)",
	)

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if erro != nil {
		return 0, nil
	}

	lastInsertedId, erro := result.LastInsertId()
	if erro != nil {
		return 0, nil
	}
}
