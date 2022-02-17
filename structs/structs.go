package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Items struct {
	gorm.Model
	ItemId      uint
	ItemCode    int64
	Description string
	Quantity    int64
	OrderId     uint
}

type Orders struct {
	gorm.Model
	OrderId      Items `gorm:"foreignkey:OrderId"`
	CustomerName string
	OrderedAt    time.Time
}
