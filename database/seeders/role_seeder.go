package seeders

import (
	"log"

	"github.com/COMTECH-63/fitness-management/models"
	"gorm.io/gorm"
)

type (
	RoleSeeder interface {
		Seed() error
		Clear() error
	}
	roleSeeder struct {
		db *gorm.DB
	}
)

func NewRoleSeeder(db *gorm.DB) RoleSeeder {
	return roleSeeder{db: db}
}

// Implement seed method
func (s roleSeeder) Seed() error {
	log.Println("RoleSeeder running...")

	// var (
	// 	role models.Role
	// )

	role := []models.Role{
		{
			Name: "Member",
		},
		{
			Name: "Employee",
		},
		{
			Name: "Trainer",
		},
	}

	result := s.db.Create(&role)
	log.Println("RoleSeeder seeded!")

	return result.Error
}

// Implement clear method
func (s roleSeeder) Clear() error {
	log.Println("Clear RoleSeeder...")
	result := s.db.Delete(&models.Role{})
	log.Println("RoleSeeder cleared!")

	return result.Error
}
