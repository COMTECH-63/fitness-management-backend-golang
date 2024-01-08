package models

type Service struct {
	Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`

	Users []User `json:"users" gorm:"many2many:service_has_users;"`
}
