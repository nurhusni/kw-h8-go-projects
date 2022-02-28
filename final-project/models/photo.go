package models

type Photo struct {
	GormModel
	Title    string
	Caption  string
	PhotoURL string
	UserID   uint
	Comments []Comment
}
