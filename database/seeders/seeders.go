package seeders

import (
	"log"

	"github.com/COMTECH-63/fitness-management/database"
	"gorm.io/gorm"
)

type seeder struct {
	userSeeder UserSeeder
}

func NewSeeder(
	db *gorm.DB,
) seeder {
	userSeeder := NewUserSeeder(db)
	return seeder{
		userSeeder: userSeeder,
	}
}

func RunSeed() {
	var err error

	database.DBConn = database.Initialize()

	seeder := NewSeeder(database.DBConn)

	// User seeder
	if err = seeder.userSeeder.Seed(); err != nil {
		log.Fatal(err)
	}
}
