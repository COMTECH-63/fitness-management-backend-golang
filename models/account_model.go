package models

type Account struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
}
