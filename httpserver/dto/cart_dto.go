package dto

type CartItemDTO struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
}

type CartDTO struct {
	ID        uint          `json:"id"`
	UserID    uint          `json:"user_id"`
	CartItems []CartItemDTO `json:"cart_items"`
	Total     float64       `json:"total"`
}
