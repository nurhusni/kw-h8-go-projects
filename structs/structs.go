package structs

import (
	"time"

	_ "gorm.io/gorm"
)

type Orders struct {
	// gorm.Model
	OrderID      uint `gorm:"primaryKey"`
	CustomerName string
	OrderedAt    time.Time
}

type Items struct {
	// gorm.Model
	ItemID      uint `gorm:"primaryKey"`
	ItemCode    int64
	Description string
	Quantity    int64
	Orders      Orders `gorm:"foreignKey:OrderID"`
}
