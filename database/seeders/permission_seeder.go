package seeders

import (
	"log"

	"github.com/COMTECH-63/fitness-management/models"
	"gorm.io/gorm"
)

type (
	PermissionSeeder interface {
		Seed() error
		Clear() error
	}
	permissionSeeder struct {
		db *gorm.DB
	}
)

func NewPermissionSeeder(db *gorm.DB) PermissionSeeder {
	return permissionSeeder{db: db}
}

// Implement seed method
func (s permissionSeeder) Seed() error {
	log.Println("PermissionSeeder running...")

	permission := []models.Permission{
		// user
		{
			Name:        "View User",
			Description: "สิทธิในการดูผู้ใช้งานในระบบ",
		},
		{
			Name:        "Create User",
			Description: "สิทธิในการสร้างผู้ใช้งานในระบบ",
		},
		{
			Name:        "Update User",
			Description: "สิทธิในการแก้ไขผู้ใช้งานในระบบ",
		},
		{
			Name:        "Delete User",
			Description: "สิทธิในการลบผู้ใช้งานในระบบ",
		},
	}

	result := s.db.Create(&permission)
	log.Println("PermissionSeeder seeded!")

	return result.Error
}

// Implement clear method
func (s permissionSeeder) Clear() error {
	log.Println("Clear PermissionSeeder...")
	result := s.db.Delete(&models.Permission{})
	log.Println("PermissionSeeder cleared!")

	return result.Error
}
