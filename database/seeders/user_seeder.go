package seeders

import (
	"log"

	"github.com/COMTECH-63/fitness-management/models"
	"golang.org/x/crypto/bcrypt"
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

	var (
		role       models.Role
		permission models.Permission
		service    models.Service
		// class                  models.Class
		// order                  models.Order
		// booking                models.Booking
		// bookingClass           models.BookingClass
		// bookingPersonalTrainer models.BookingPersonalTrainer
	)

	users := []models.User{
		{
			Username:    "admin",
			Password:    "passw0rd",
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
			Username:    "karn001",
			Password:    "karnpassword1",
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
		{

			Username:    "peempot002",
			Password:    "peempotpassword2",
			FirstName:   "ภีมพศ",
			LastName:    "ภาดาสิทธิภูมิ",
			IDCard:      "111111112-3",
			Email:       "peempot_psp@fitness.co.th",
			PhoneNumber: "0111111112",
			Address:     "ปทุมธานี คลอง 6",
			Sex:         models.Male,
			ImageURL:    "https://cdn.example.com/user/3/avatar.png",
			MemberID:    "00002",
		},
		{
			Username:    "thanachok003",
			Password:    "thanachokpassword03",
			FirstName:   "ธนโชค",
			LastName:    "จอมคำสิงห์",
			IDCard:      "111111112-4",
			Email:       "thanachok_jks@fitness.co.th",
			PhoneNumber: "0111111113",
			Address:     "สายไหม",
			Sex:         models.Male,
			ImageURL:    "https://cdn.example.com/user/4/avatar.png",
			MemberID:    "00003",
		},
	}

	// Hash passwords for all users
	for i := range users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(users[i].Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		users[i].Password = string(hashedPassword)
	}

	result := s.db.Create(&users)

	s.db.Find(&role)
	s.db.Find(&permission)
	s.db.Find(&service)
	// s.db.Find(class)
	// s.db.Find(order)
	// s.db.Find(booking)
	// s.db.Find(bookingClass)
	// s.db.Find(bookingPersonalTrainer)

	s.db.Model(&users).Association("Roles").Append(&role)
	s.db.Model(&users).Association("Permissions").Append(&permission)
	s.db.Model(&users).Association("Services").Append(&service)
	// s.db.Model(&users[i]).Association("Classes").Append(&class)
	// s.db.Model(&users[i]).Association("Orders").Append(&order)
	// s.db.Model(&users[i]).Association("Bookings").Append(&booking)
	// s.db.Model(&users[i]).Association("BookingClasses").Append(&bookingClass)
	// s.db.Model(&users[i]).Association("BookingPersonalTrainers").Append(&bookingPersonalTrainer)

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
