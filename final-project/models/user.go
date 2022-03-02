package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username     string        `gorm:"not null;unique" json:"username" form:"username" valid:"required~Username is required"`
	Email        string        `gorm:"not null;unique" json:"email" form:"email" valid:"required~Email is required,email~Invalid email format"`
	Password     string        `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age          uint          `gorm:"not null" json:"age" form:"age" valid:"required~Age is required,range(8|100)~Minimum age is 8 years old"`
	Photos       []Photo       `json:"photos,omitempty"`
	Comments     []Comment     `json:"comments,omitempty"`
	SocialMedias []SocialMedia `json:"social_medias,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
