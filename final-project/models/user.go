package models

type User struct {
	GormModel
	Username     string        `gorm:"not null" json:"username"`
	Email        string        `gorm:"not null" json:"email"`
	Password     string        `gorm:"not null" json:"password"`
	Age          uint          `gorm:"not null" json:"age"`
	Photos       []Photo       `json:"-"`
	Comments     []Comment     `json:"-"`
	SocialMedias []SocialMedia `json:"-"`
}
