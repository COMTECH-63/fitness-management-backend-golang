package seeders

import (
	"log"

	"github.com/COMTECH-63/fitness-management/models"
	"gorm.io/gorm"
)

type (
	UserSeeder interface {
		Seed() error
		Clear() error
	}
	userSeeder struct {
		db *gorm.DB
	}
)

func NewUserSeeder(db *gorm.DB) UserSeeder {
	return userSeeder{db: db}
}

// Implement seed method
func (s userSeeder) Seed() error {
	log.Println("User Seeder running...")

	user := []models.User{
		{
			FirstName:   "Super",
			LastName:    "Administrator",
			IDCard:      "-",
			Email:       "superadmin@fitness.co.th",
			PhoneNumber: "-",
			Address:     "-",
			Sex:         models.Male,
			ImageURL:    "https://cdn.example.com/user/1/avatar.png",
			MemberID:    "00000",
		},
		{
			FirstName:   "กานต์",
			LastName:    "ลพสุนทร",
			IDCard:      "111111112-1",
			Email:       "karn_lst@fitness.co.th",
			PhoneNumber: "0111111111",
			Address:     "รามอินทรา",
			Sex:         models.Male,
			ImageURL:    "https://cdn.example.com/user/2/avatar.png",
			MemberID:    "00001",
		},
	}

	result := s.db.Create(&user)
	log.Println("User Seeder seeded!")

	return result.Error
}

// Implement clear method
func (s userSeeder) Clear() error {
	log.Println("Clear UserSeeder...")
	result := s.db.Delete(&models.User{})
	log.Println("UserSeeder cleared!")

	return result.Error
}
