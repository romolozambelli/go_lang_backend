package repo

import (
	"backend/src/models"
	"database/sql"
	"fmt"
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
		"INSERT INTO users (name,nickname,email,password) values (?,?,?,?)",
	)

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if erro != nil {
		return 0, nil
	}

	lastInsertedID, erro := result.LastInsertId()
	if erro != nil {
		return 0, nil
	}

	fmt.Printf("User inserted with success on repo = %d/n", lastInsertedID)

	return uint64(lastInsertedID), nil
}

func (repoUser Users) GetUserOrNick(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) //%nameOrNick%

	lines, erro := repoUser.db.Query(
		"SELECT id, name, nickname, email, created FROM users WHERE name LIKE ? OR nickname LIKE ?",
		nameOrNick, nameOrNick)

	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreateDate,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)

	}
	return users, nil
}
