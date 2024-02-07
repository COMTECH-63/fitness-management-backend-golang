package models

type OrderItem struct {
	Model
	OrderID uint  `json:"order_id"`
	Order   Order `json:"order" gorm:"foreignkey:OrderID"`

	ServiceID uint    `json:"service_id"`
	Service   Service `json:"service" gorm:"foreignkey:ServiceID"`
}
