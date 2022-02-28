package models

import "time"

type GormModel struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
