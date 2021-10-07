package main

import (
	"fmt"
	"strconv"
)

func main() {
	// soal 1

	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	p, err := strconv.Atoi(panjangPersegiPanjang)
	l, err := strconv.Atoi(lebarPersegiPanjang)
	a, err := strconv.Atoi(alasSegitiga)
	t, err := strconv.Atoi(tinggiSegitiga)

	if err == nil {
		var kelilingPersegiPanjang int = 2 * (p + l)
		var luasSegitiga int = a * t / 2

		fmt.Println("-----SOAL 1-----")
		fmt.Println("Keliling Persegi Panjang :", kelilingPersegiPanjang)
		fmt.Println("Luas Segitiga :", luasSegitiga)
	}

	// soal 2
	fmt.Println("-----SOAL 2-----")
	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiJohn >= 80 {
		fmt.Println("Nilai John A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("Nilai John B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("Nilai John C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("Nilai John D")
	} else {
		fmt.Println("Nilai John E")
	}

	if nilaiDoe >= 80 {
		fmt.Println("Nilai John A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("Nilai John B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("Nilai John C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("Nilai John D")
	} else {
		fmt.Println("Nilai John E")
	}

	// soal 3
	fmt.Println("-----SOAL 3-----")

	var tanggal = 17
	var bulan = 8
	var tahun = 2000

	switch bulan {
	case 1:
		fmt.Println(tanggal, "Januari", tahun)
	case 2:
		fmt.Println(tanggal, "Februari", tahun)
	case 3:
		fmt.Println(tanggal, "Maret", tahun)
	case 4:
		fmt.Println(tanggal, "April", tahun)
	case 5:
		fmt.Println(tanggal, "Mei", tahun)
	case 6:
		fmt.Println(tanggal, "Juni", tahun)
	case 7:
		fmt.Println(tanggal, "Juli", tahun)
	case 8:
		fmt.Println(tanggal, "Agustus", tahun)
	case 9:
		fmt.Println(tanggal, "September", tahun)
	case 10:
		fmt.Println(tanggal, "Oktober", tahun)
	case 11:
		fmt.Println(tanggal, "November", tahun)
	case 12:
		fmt.Println(tanggal, "Desember", tahun)
	}

	// 	aby boomer, kelahiran 1944 s.d 1964
	// Generasi X, kelahiran 1965 s.d 1979
	// Generasi Y (Millenials), kelahiran 1980 s.d 1994
	// Generasi Z, kelahiran 1995 s.d 2015
	// soal 4
	fmt.Println("-----SOAL 4-----")
	switch {
	case tahun >= 1944 && tahun <= 1964:
		fmt.Println("Baby Boomer")
	case tahun >= 1965 && tahun <= 1979:
		fmt.Println("Generasi X")
	case tahun >= 1980 && tahun <= 1994:
		fmt.Println("Generasi Y (Millenials)")
	case tahun >= 1995 && tahun <= 2015:
		fmt.Println("Generasi Z")
	}
}
