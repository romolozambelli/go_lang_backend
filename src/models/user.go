package models

import (
	"backend/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID         uint64    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Nickname   string    `json:"nickname,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	CreateDate time.Time `json:"created,omitempty"`
}

// Prepare the user with the validations and cleanning
func (user *User) Prepare(step string) error {
	if erro := user.validateUser(step); erro != nil {
		return erro
	}
	if erro := user.format(step); erro != nil {
		return erro
	}
	return nil
}

func (user *User) validateUser(step string) error {
	if user.Name == "" {
		return errors.New("name is mandatory to create a user")

	}

	if user.Nickname == "" {
		return errors.New("nickname is mandatory to create a user")

	}

	if user.Email == "" {
		return errors.New("email is mandatory to create a user")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return checkmail.ErrBadFormat
	}

	if step == "register" && user.Password == "" {
		return errors.New("password is mandatory to create a user")

	}
	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		passwordWithHash, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}
		user.Password = string(passwordWithHash)
	}
	return nil
}
