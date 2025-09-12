package models

type Order struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	Amount uint
	UserID *uint `json:"user_id"` // nullable и тег для JSON
	User   User  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
