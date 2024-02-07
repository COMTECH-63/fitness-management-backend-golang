package models

import "time"

type Booking struct {
	Model
	Date time.Time `json:"date"`

	UserID uint `json:"user_id"`
	User   User `json:"user" gorm:"foreignkey:UserID"`
}
