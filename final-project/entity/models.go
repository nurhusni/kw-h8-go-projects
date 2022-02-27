package entity

import "time"

type User struct {
	ID           uint `gorm:"primaryKey"`
	Username     string
	Email        string
	Password     string
	Age          uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Photos       []Photo
	Comments     []Comment
	SocialMedias []SocialMedia
}

type Photo struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Caption   string
	PhotoURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint
	Comments  []Comment
	// UserID (foreign key of User)
}

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint
	PhotoID   uint
	// UserID (foreign key of User)
	// PhotoID (foreign key of Photo)
}

type SocialMedia struct {
	ID             uint `gorm:"primaryKey"`
	Name           string
	SocialMediaURL string
	UserID         uint
	// UserID
}
