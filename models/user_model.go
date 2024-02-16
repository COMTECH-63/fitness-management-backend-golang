package models

type Sex string

const (
	Male   Sex = "male"
	Female Sex = "female"
	O
)

func (u *User) SetSex(value string) {
	if value == "male" {
		u.Sex = Male
	} else if value == "female" {
		u.Sex = Female
	}
}

type User struct {
	Model
	Username    string `json:"username"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	IDCard      string `json:"id_card"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Sex         Sex    `json:"sex"`
	ImageURL    string `json:"image_url"`
	MemberID    string `json:"member_id"`

	Roles                   []Role                   `json:"roles" gorm:"many2many:role_has_users;"`
	Permissions             []Permission             `json:"permissions" gorm:"many2many:user_has_permissions;"`
	Services                []Service                `json:"services" gorm:"many2many:service_has_users;"`
	Classes                 []Class                  `json:"classes" gorm:"many2many:class_has_users;"`
	Orders                  []Order                  `json:"orders" gorm:"foreignkey:UserID"`
	Bookings                []Booking                `json:"bookings" gorm:"foreignkey:UserID"`
	BookingClasses          []BookingClass           `json:"booking_classes" gorm:"foreignkey:UserID"`
	BookingPersonalTrainers []BookingPersonalTrainer `json:"booking_personal_trainers" gorm:"foreignkey:UserID"`
}
