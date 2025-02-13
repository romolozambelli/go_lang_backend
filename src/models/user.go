package models

import (
	"errors"
	"strings"
	"time"
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
func (user *User) Prepare() error {
	if erro := user.validateUser(); erro != nil {
		return erro
	}
	user.format()
	return nil
}

func (user *User) validateUser() error {

	if user.Name == "" {
		return errors.New("name is mandatory to create a user")

	}

	if user.Nickname == "" {
		return errors.New("nickname is mandatory to create a user")

	}

	if user.Email == "" {
		return errors.New("email is mandatory to create a user")

	}

	if user.Password == "" {
		return errors.New("password is mandatory to create a user")

	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)

}
