package models

type User struct {
	GormModel
	Username     string        `json:"username"`
	Email        string        `json:"email"`
	Password     string        `json:"password"`
	Age          uint          `json:"age"`
	Photos       []Photo       `json:"-"`
	Comments     []Comment     `json:"-"`
	SocialMedias []SocialMedia `json:"-"`
}
