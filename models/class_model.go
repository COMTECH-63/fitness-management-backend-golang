package models

type MotivationC string
type Intensity string
type MinuteC string

const (
	LeanC       MotivationC = "lean"
	PerformaceC MotivationC = "performance"
	StrongC     MotivationC = "strong"
	WellBeingC  MotivationC = "well_being"
	LowC        Intensity   = "low"
	MidC        Intensity   = "mid"
	HighC       Intensity   = "high"
	Minute30C   MinuteC     = "30"
	Minute45C   MinuteC     = "45"
	Minute60C   MinuteC     = "60"
	Minute95C   MinuteC     = "95"
)

func (c *Class) SetMotivationC(value string) {
	if value == "lean" {
		c.MotivationC = LeanC
	} else if value == "performance" {
		c.MotivationC = PerformaceC
	} else if value == "strong" {
		c.MotivationC = StrongC
	} else if value == "well_being" {
		c.MotivationC = WellBeingC
	}
}

func (c *Class) SetIntensity(value string) {
	if value == "low" {
		c.Intensity = LowC
	} else if value == "mid" {
		c.Intensity = MidC
	} else if value == "high" {
		c.Intensity = HighC
	}
}

func (c *Class) SetMinuteC(value string) {
	if value == "30" {
		c.MinuteC = Minute30C
	} else if value == "45" {
		c.MinuteC = Minute45C
	} else if value == "60" {
		c.MinuteC = Minute60C
	} else if value == "95" {
		c.MinuteC = Minute95C
	}
}

type Class struct {
	Model
	Name        string      `json:"name"`
	Description string      `json:"description"`
	MotivationC MotivationC `json:"motivation"`
	Intensity   Intensity   `json:"intensity"`
	MinuteC     MinuteC     `json:"minute"`

	UserID uint `json:"user_id"`
	User   User `json:"user" gorm:"foreignkey:UserID"`

	Users []User `json:"users" gorm:"many2many:class_has_users;"`
}
