package main

import (
	"net/http"

	"go-assignment-3/controllers"
)

var PORT = ":3000"

func main() {
	http.HandleFunc("/", controllers.ShowStatus)
	http.ListenAndServe(PORT, nil)
}
