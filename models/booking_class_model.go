package models

import "time"

type BookingClass struct {
	Model
	Date time.Time `json:"date"`

	ClassID uint  `json:"class_id"`
	Class   Class `json:"class" gorm:"foreignkey:ClassID"`
	UserID  uint  `json:"user_id"`
	User    User  `json:"user" gorm:"foreignkey:UserID"`
}
