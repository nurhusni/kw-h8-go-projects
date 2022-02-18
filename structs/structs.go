package structs

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerName string
	OrderedAt    time.Time
	Item         Item `gorm:"foreignKey:OrderID;references:ID"`
}

type Item struct {
	gorm.Model
	ItemCode    int64
	Description string
	Quantity    int64
	OrderID     uint
}
