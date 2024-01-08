package models

type MotivationP string
type MinuteP string

const (
	LeanP       MotivationP = "lean"
	PerformaceP MotivationP = "performance"
	StrongP     MotivationP = "strong"
	WellBeingP  MotivationP = "well_being"
	Minute30P   MinuteP     = "30"
	Minute45P   MinuteP     = "45"
	Minute60P   MinuteP     = "60"
	Minute95P   MinuteP     = "95"
)

func (p *PersonalTrainer) SetMotivationP(value string) {
	if value == "lean" {
		p.MotivationP = LeanP
	} else if value == "performance" {
		p.MotivationP = PerformaceP
	} else if value == "strong" {
		p.MotivationP = StrongP
	} else if value == "well_being" {
		p.MotivationP = WellBeingP
	}
}

func (p *PersonalTrainer) SetMinuteP(value string) {
	if value == "30" {
		p.MinuteP = Minute30P
	} else if value == "45" {
		p.MinuteP = Minute45P
	} else if value == "60" {
		p.MinuteP = Minute60P
	} else if value == "95" {
		p.MinuteP = Minute95P
	}
}

type PersonalTrainer struct {
	Model
	Name        string      `json:"name"`
	Description string      `json:"description"`
	MotivationP MotivationP `json:"motivation"`
	Intensity   Intensity   `json:"intensity"`
	MinuteP     MinuteP     `json:"minute"`

	UserID uint `json:"user_id"`
	User   User `json:"user" gorm:"foreignkey:UserID"`
}
