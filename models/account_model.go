package models

type Account struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`

	UserID uint `json:"user_id"`
	User   User `json:"user" gorm:"foreignkey:UserID"`
}
