package models

import "time"

type Status string

const (
	NotPaid Status = "not_paid"
	Paid    Status = "paid"
)

func (o *OrderPayment) SetStatus(value string) {
	if value == "not_paid" {
		o.Status = NotPaid
	} else if value == "paid" {
		o.Status = Paid
	}
}

type OrderPayment struct {
	Model
	Amount float64   `json:"amount"`
	Status Status    `json:"status"`
	Date   time.Time `json:"Date"`

	OrderID uint  `json:"order_id"`
	Order   Order `json:"order" gorm:"foreignkey:OrderID"`
}
