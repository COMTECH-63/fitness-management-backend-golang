package models

type Order struct {
	Model
	Total    float64 `json:"total"`
	Vat      float64 `json:"vat"`
	TotalVat float64 `json:"total_vat"`

	UserID uint `json:"user_id"`
	User   User `json:"user" gorm:"foreignkey:UserID"`

	OrderItems []OrderItem `json:"order_items" gorm:"foreignkey:OrderID"`
}
