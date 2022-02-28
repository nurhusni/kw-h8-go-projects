package models

type SocialMedia struct {
	ID             uint `gorm:"primaryKey"`
	Name           string
	SocialMediaURL string
	UserID         uint
}
