package models

import "time"

type BookingPersonalTrainer struct {
	Model
	Date time.Time `json:"date"`

	PersonalTrainerID uint            `json:"personal_trainer_id"`
	PersonalTrainer   PersonalTrainer `json:"personal_trainer" gorm:"foreignkey:PersonalTrainerID"`
	UserID            uint            `json:"user_id"`
	User              User            `json:"user" gorm:"foreignkey:UserID"`
}
