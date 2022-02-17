package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Orders struct {
	gorm.Model
	// OrderId      uint `gorm:"primarykey"`
	CustomerName string
	OrderedAt    time.Time
	Items        Items
}

type Items struct {
	gorm.Model
	// ItemId      uint `gorm:"primarykey"`
	ItemCode    int64
	Description string
	Quantity    int64
	// OrderId     uint `gorm:"foreignKey:ID"`
}
