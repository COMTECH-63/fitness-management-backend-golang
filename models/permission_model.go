package models

type Permission struct {
	Model
	Description string `json:"description"`

	Roles []Role `json:"roles" gorm:"many2many:role_has_permissions;"`
	Users []User `json:"users" gorm:"many2many:user_has_permissions;"`
}