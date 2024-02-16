package seeders

import (
	"log"

	"github.com/COMTECH-63/fitness-management/database"
	"gorm.io/gorm"
)

type seeder struct {
	userSeeder       UserSeeder
	roleSeeder       RoleSeeder
	permissionSeeder PermissionSeeder
	serviceSeeder    ServiceSeeder
}

func NewSeeder(
	db *gorm.DB,
) seeder {

	userSeeder := NewUserSeeder(db)
	roleSeeder := NewRoleSeeder(db)
	permissionSeeder := NewPermissionSeeder(db)
	serviceSeeder := NewServiceSeeder(db)

	return seeder{
		userSeeder:       userSeeder,
		roleSeeder:       roleSeeder,
		permissionSeeder: permissionSeeder,
		serviceSeeder:    serviceSeeder,
	}
}

func RunSeed() {
	var err error

	database.DBConn = database.Initialize()

	seeder := NewSeeder(database.DBConn)

	// Role seeder
	if err = seeder.roleSeeder.Seed(); err != nil {
		log.Fatal(err)
	}

	// Permission seeder
	if err = seeder.permissionSeeder.Seed(); err != nil {
		log.Fatal(err)
	}

	// User seeder
	if err = seeder.userSeeder.Seed(); err != nil {
		log.Fatal(err)
	}

	// Service seeder
	if err = seeder.serviceSeeder.Seed(); err != nil {
		log.Fatal(err)
	}
}
