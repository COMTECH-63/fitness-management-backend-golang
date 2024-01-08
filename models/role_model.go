package models

type Role struct {
	Model
	Name string `json:"name"`

	Users       []User       `json:"users" gorm:"many2many:role_has_users;"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_has_permissions;"`
}
