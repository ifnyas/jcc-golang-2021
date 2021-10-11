package util

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, p interface{}, status int) {
	encoded, err := json.Marshal(p)
	if err != nil {
		http.Error(w, "Oops...", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(encoded))
}

func ErrHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LocBetween(lat float64, lng float64, radius float64) []float64 {
	earth := 6378.137
	m := (1 / ((2 * math.Pi / 360) * earth)) / 1000

	rLat := radius * m
	minLat := lat - rLat
	maxLat := lat + rLat

	rLng := rLat / math.Cos(lat*(math.Pi/180))
	minLng := lng - rLng
	maxLng := lng + rLng

	return []float64{minLat, maxLat, minLng, maxLng}
}
