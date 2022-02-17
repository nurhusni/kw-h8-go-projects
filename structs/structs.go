package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Items struct {
	gorm.Model
	ItemId      int
	ItemCode    int
	Description string
	Quantity    int
	OrderId     int
}

type Orders struct {
	gorm.Model
	OrderId      int
	CustomerName string
	OrderedAt    time.Time
}
