package seeders

import (
	"log"

	"github.com/COMTECH-63/fitness-management/models"
	"gorm.io/gorm"
)

type (
	ServiceSeeder interface {
		Seed() error
		Clear() error
	}
	serviceSeeder struct {
		db *gorm.DB
	}
)

func NewServiceSeeder(db *gorm.DB) ServiceSeeder {
	return serviceSeeder{db: db}
}

// Implement seed method
func (s serviceSeeder) Seed() error {
	log.Println("ServiceSeeder running...")

	var user models.User

	services := []models.Service{
		{
			Name:        "Membership 1 month",
			Description: "สมาชิกรายเดือน 1 เดือน",
			Price:       1499,
		},
		{
			Name:        "Membership 3 month",
			Description: "สมาชิกรายเดือน 3 เดือน",
			Price:       3499,
		},
		{
			Name:        "Membership 6 month",
			Description: "สมาชิกรายเดือน 6 เดือน",
			Price:       5499,
		},
		{
			Name:        "Membership 9 month",
			Description: "สมาชิกรายเดือน 1 เดือน",
			Price:       7499,
		},
		{
			Name:        "Membership 12 month",
			Description: "สมาชิกรายเดือน 12 เดือน",
			Price:       9499,
		},
		{
			Name:        "Membership + Trainer 1 month",
			Description: "สมาชิกรายเดือน + เทรนเนอร์ 1 เดือน",
			Price:       3499,
		},
		{
			Name:        "Membership + Trainer 3 month",
			Description: "สมาชิกรายเดือน + เทรนเนอร์ 3 เดือน",
			Price:       5499,
		},
		{
			Name:        "Membership + Trainer 6 month",
			Description: "สมาชิกรายเดือน + เทรนเนอร์ 6 เดือน",
			Price:       7499,
		},
		{
			Name:        "Membership + Trainer 9 month",
			Description: "สมาชิกรายเดือน + เทรนเนอร์ 1 เดือน",
			Price:       9499,
		},
		{
			Name:        "Membership + Trainer 12 month",
			Description: "สมาชิกรายเดือน + เทรนเนอร์ 12 เดือน",
			Price:       12499,
		},
	}

	result := s.db.Create(&services)

	s.db.Find(&user)

	s.db.Model(&services).Association("Users").Append(&user)

	log.Println("ServiceSeeder seeded!")

	return result.Error
}

// Implement clear method
func (s serviceSeeder) Clear() error {
	log.Println("Clear ServiceSeeder...")
	result := s.db.Delete(&models.Service{})
	log.Println("ServiceSeeder cleared!")

	return result.Error
}
