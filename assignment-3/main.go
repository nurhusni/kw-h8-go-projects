package main

import (
	"net/http"

	"go-assignment-3/controllers"
)

var PORT = ":3000"

func init() {
	values := controllers.GenerateStatus()
	_ = values
}

func main() {
	http.HandleFunc("/", controllers.ShowStatus)
	http.ListenAndServe(PORT, nil)
}
