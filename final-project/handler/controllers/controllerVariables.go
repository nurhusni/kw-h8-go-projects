package controllers

import "khg-final-project/infra"

var (
	appJson = "application/json"
	db      = infra.GetDB()
	err     error
)
