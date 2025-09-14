package orders

type OrderCreatedEvent struct {
	OrderID uint   `json:"order_id"`
	UserID  uint   `json:"user_id"`
	Status  string `json:"status"`
}
