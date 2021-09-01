package main

import (
	"fmt"
	"strconv"
)

func main() {
	// soal 1
	panjangPersegiPanjang := "8"
	lebarPersegiPanjang := "5"

	alasSegitiga := "6"
	tinggiSegitiga := "7"

	pppInt, _ := strconv.Atoi(panjangPersegiPanjang)
	lppInt, _ := strconv.Atoi(lebarPersegiPanjang)
	asInt, _ := strconv.Atoi(alasSegitiga)
	tsInt, _ := strconv.Atoi(tinggiSegitiga)

	kelilingPersegiPanjang := (pppInt + lppInt) * 2
	luasSegitiga := asInt * tsInt / 2

	fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang, "; Luas Segitiga:", luasSegitiga)

	// soal 2
	nilaiJohn := 80
	nilaiDoe := 50

	indeksJohn := ""
	switch {
	case nilaiJohn >= 80:
		indeksJohn = "A"
	case nilaiJohn >= 70:
		indeksJohn = "B"
	case nilaiJohn >= 60:
		indeksJohn = "C"
	case nilaiJohn >= 50:
		indeksJohn = "D"
	default:
		indeksJohn = "E"
	}

	indeksDoe := ""
	switch {
	case nilaiDoe >= 80:
		indeksDoe = "A"
	case nilaiDoe >= 70:
		indeksDoe = "B"
	case nilaiDoe >= 60:
		indeksDoe = "C"
	case nilaiDoe >= 50:
		indeksDoe = "D"
	default:
		indeksDoe = "E"
	}

	fmt.Println("Indeks John:", indeksJohn, "; Indeks Doe:", indeksDoe)

	// soal 3
	tanggal := 2
	bulan := 8
	tahun := 1995

	bulanStr := ""
	switch bulan {
	case 1:
		bulanStr = "Januari"
	case 2:
		bulanStr = "Februari"
	case 3:
		bulanStr = "Maret"
	case 4:
		bulanStr = "April"
	case 5:
		bulanStr = "Mei"
	case 6:
		bulanStr = "Juni"
	case 7:
		bulanStr = "Juli"
	case 8:
		bulanStr = "Agustus"
	case 9:
		bulanStr = "September"
	case 10:
		bulanStr = "Oktober"
	case 11:
		bulanStr = "November"
	case 12:
		bulanStr = "Desember"
	}

	dateStr := strconv.Itoa(tanggal) + " " + bulanStr + " " + strconv.Itoa(tahun)
	fmt.Println(dateStr)

	// soal 4
	generasi := ""
	switch {
	case tahun >= 1995 && tahun <= 2015:
		generasi = "Generasi Z"
	case tahun >= 1980 && tahun <= 1994:
		generasi = "Generasi Y (Millenials)"
	case tahun >= 1965 && tahun <= 1979:
		generasi = "Generasi X"
	case tahun >= 1944 && tahun <= 1964:
		generasi = "Baby boomer"
	default:
		generasi = "Tidak diketahui"
	}
	fmt.Println(generasi)
}
