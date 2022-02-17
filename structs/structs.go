package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Items struct {
	gorm.Model
	ItemId      uint
	ItemCode    uint
	Description string
	Quantity    uint
	OrderId     Orders `gorm:"foreignkey:OrderId"`
}

type Orders struct {
	gorm.Model
	OrderId      uint
	CustomerName string
	OrderedAt    time.Time
	Items        []Items
}
