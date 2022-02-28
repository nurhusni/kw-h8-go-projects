package models

type Comment struct {
	GormModel
	Message string
	UserID  uint
	PhotoID uint
}
