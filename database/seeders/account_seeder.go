package seeders

import (
	"log"

	"github.com/COMTECH-63/fitness-management/models"
	"gorm.io/gorm"
)

type (
	AccountSeeder interface {
		Seed() error
		Clear() error
	}
	accountSeeder struct {
		db *gorm.DB
	}
)

func NewAccountSeeder(db *gorm.DB) AccountSeeder {
	return accountSeeder{db: db}
}

// Implement seed method
func (s accountSeeder) Seed() error {
	log.Println("Account Seeder running...")

	var user models.User

	accounts := []models.Account{
		{
			Username: "admin",
			Password: "passw0rd",
		},
		{
			Username: "karn001",
			Password: "karnpassword1",
		},
		{
			Username: "peempot002",
			Password: "peempotpassword2",
		},
		{
			Username: "thanachok003",
			Password: "thanachokpassword03",
		},
	}

	result := s.db.Create(&accounts)

	s.db.Find(&user)

	log.Println("Account Seeder seeded!")

	return result.Error
}

// Implement clear method
func (s accountSeeder) Clear() error {
	log.Println("Clear AccountSeeder...")
	result := s.db.Delete(&models.Account{})
	log.Println("AccountSeeder cleared!")

	return result.Error
}
