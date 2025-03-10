package repo

import (
	"backend/src/models"
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
		"INSERT INTO users (name,nickname,email,password) VALUES (?,?,?,?)",
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

	return uint64(lastInsertedID), nil
}

func (repoUser Users) GetUserOrNick(nameOrNick string) ([]models.User, error) {

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

// Get a user based on a given id from the datbase
func (repoUser Users) SearchUserByID(ID uint64) (models.User, error) {

	lines, erro := repoUser.db.Query(
		"SELECT id, name, nickname, email, created FROM users WHERE id = ?",
		ID)

	if erro != nil {
		return models.User{}, erro
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreateDate,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil

}

// Update the user data with a given ID
func (repoUser Users) UpdateUser(ID uint64, user models.User) error {

	statement, erro := repoUser.db.Prepare(
		"UPDATE users SET name = ?, nickname = ?, email=? WHERE id = ?",
	)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(user.Name, user.Nickname, user.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// Delete a user data with a given ID
func (repoUser Users) DeleteUserByID(ID uint64) error {

	statement, erro := repoUser.db.Prepare(
		"DELETE FROM users WHERE id = ?",
	)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// Search a user by email on the DB and return id and password
func (repoUser Users) SearchByEmail(email string) (models.User, error) {

	line, erro := repoUser.db.Query(
		"SELECT id, password FROM users WHERE email = ?", email)

	if erro != nil {
		return models.User{}, erro
	}
	defer line.Close()

	var user models.User

	for line.Next() {
		if erro = line.Scan(&user.ID, &user.Password); erro != nil {
			return models.User{}, erro
		}
	}

	return user, erro
}

func (repoUser Users) Follow(userID, followID uint64) error {
	statement, erro := repoUser.db.Prepare(
		"INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?,?)",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID, followID); erro != nil {
		return erro
	}
	return nil
}

func (repoUser Users) UnFollow(userID, followID uint64) error {

	statement, erro := repoUser.db.Prepare(
		"DELETE FROM followers WHERE user_id = ? AND follower_id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID, followID); erro != nil {
		return erro
	}
	return nil
}

// Get all the followers from a giver user
func (repoUser Users) GetFollowers(userID uint64) ([]models.User, error) {

	lines, erro := repoUser.db.Query(
		`SELECT u.id, u.name, u.nickname, u.email,u.created
		FROM users u INNER JOIN followers s on u.id = s.follower_id 
		WHERE s.user_id = ?`, userID)

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

// Get all the users is follow by the given users
func (repoUser Users) GetFollowingUsers(userID uint64) ([]models.User, error) {

	lines, erro := repoUser.db.Query(
		`SELECT u.id, u.name, u.nickname, u.email,u.created
		FROM users u INNER JOIN followers s on u.id = s.user_id 
		WHERE s.follower_id = ?`, userID)

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

// Get the user password
func (repoUser Users) GetPassword(userID uint64) (string, error) {

	line, erro := repoUser.db.Query(
		"SELECT password FROM users WHERE id = ?", userID)

	if erro != nil {
		return "", erro
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if erro = line.Scan(&user.Password); erro != nil {
			return "", erro
		}
	}

	return user.Password, nil

}

func (repoUser Users) UpdatePassword(userID uint64, password string) error {

	statement, erro := repoUser.db.Prepare(
		"UPDATE users SET password = ? WHERE id = ?",
	)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(password, userID); erro != nil {
		return erro
	}

	return nil

}
