package models

type Order struct {
	ID     uint        `gorm:"primaryKey" json:"id"`
	Name   string      `json:"name"`
	Amount uint        `json:"amount"`
	UserID *uint       `json:"user_id"` // nullable
	User   User        `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	Items  []OrderItem `gorm:"foreignKey:OrderID" json:"items,omitempty"`
}
