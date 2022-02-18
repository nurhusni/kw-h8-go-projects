package structs

import (
	"time"

	"gorm.io/gorm"
	// _ "gorm.io/gorm"
)

type Order struct {
	gorm.Model
	// OrderID      uint `gorm:"primaryKey;autoIncrement:true"`
	CustomerName string
	OrderedAt    time.Time
	Item         Item `gorm:"foreignKey:OrderID;references:ID"`
	// Item Item `gorm:"foreignKey:OrderID;references:OrderID"`
}

type Item struct {
	gorm.Model
	// ItemID      uint `gorm:"primaryKey;autoIncrement:true"`
	ItemCode    int64
	Description string
	Quantity    int64
	OrderID     uint
}
