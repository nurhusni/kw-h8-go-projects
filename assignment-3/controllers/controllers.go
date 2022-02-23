package controllers

import (
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
)

type Values struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	WaterStatus string `json:"waterStatus"`
	WindStatus  string `json:"windStatus"`
}

// type status interface {
// 	getStatus() string
// }

type wind struct {
	value int
}

type water struct {
	value int
}

func (wa water) getStatus() string {
	var status string

	switch {
	case wa.value <= 5:
		status = "aman"
	case wa.value >= 6 && wa.value <= 8:
		status = "siaga"
	default:
		status = "bahaya"
	}

	return status
}

func (wi wind) getStatus() string {
	var status string

	switch {
	case wi.value <= 6:
		status = "aman"
	case wi.value >= 7 && wi.value <= 15:
		status = "siaga"
	default:
		status = "bahaya"
	}

	return status
}

func getData() (int, int) {
	water := rand.Intn(100)
	wind := rand.Intn(100)

	return water, wind
}

func ShowStatus(w http.ResponseWriter, r *http.Request) {
	waterValue, windValue := getData()

	values := &Values{
		Water: waterValue,
		Wind:  windValue,
	}

	wa := water{value: values.Water}
	wi := wind{value: values.Wind}

	waterStatus := wa.getStatus()
	windStatus := wi.getStatus()

	_, _ = waterStatus, windStatus

	v, err := json.Marshal(values)
	if err != nil {
		panic(err)
	}

	if r.Method == "GET" {
		tpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, waterStatus)
		tpl.Execute(w, windStatus)

		json.NewEncoder(w).Encode(v)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}
