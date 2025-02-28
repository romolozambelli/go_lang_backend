package models

type Password struct {
	Password    string `json:"password,omitempty"`
	NewPassword string `json:"newpassword,omitempty"`
}
