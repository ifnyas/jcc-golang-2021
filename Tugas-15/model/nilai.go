package model

import "time"

type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	IndeksNilai string `json:"indeks_nilai"`
	Nilai, ID   uint
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetIndeks(n uint) string {
	indeks := ""
	switch {
	case n >= 80:
		indeks = "A"
	case n >= 70:
		indeks = "B"
	case n >= 60:
		indeks = "C"
	case n >= 50:
		indeks = "D"
	default:
		indeks = "E"
	}
	return indeks
}
