package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	var luasLingkaran float64
	var kelilingLingkaran float64

	circleCount := func(radius *float64) {
		luasLingkaran = 3.14 * math.Pow(*radius, 2)
		kelilingLingkaran = 2 * 3.14 * *radius
	}

	var radius float64 = 4
	var rPointer *float64 = &radius

	circleCount(rPointer)
	fmt.Println("Jari2 =", radius, "; Luas =", luasLingkaran, "; Keliling =", kelilingLingkaran)

	// soal 2
	var sentence string
	introduce := func(sent *string, name string, gender string, job string, age string) {
		title := ""
		if gender == "laki-laki" {
			title = "Pak"
		} else {
			title = "Bu"
		}
		*sent = title + " " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	}

	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)

	// soal 3
	var buah = []string{}
	var bPointer = &buah
	*bPointer = append(*bPointer, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")
	for i, item := range *bPointer {
		fmt.Println(strconv.Itoa(i+1) + ". " + item)
	}

	// soal 4
	var dataFilm = []map[string]string{}
	tambahDataFilm := func(title string, jam string, genre string, tahun string, dfPointer *[]map[string]string) {
		film := map[string]string{
			"title":    title,
			"duration": jam,
			"genre":    genre,
			"year":     tahun,
		}
		*dfPointer = append(*dfPointer, film)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	// isi dengan jawaban anda untuk menampilkan data
	for i, item := range dataFilm {
		indent := math.Floor(float64(i)/10) + 2
		fmt.Println(strconv.Itoa(i+1) + ". title : " + item["title"])
		fmt.Println(strings.Repeat(" ", int(indent)) + " duration : " + item["duration"])
		fmt.Println(strings.Repeat(" ", int(indent)) + " genre : " + item["genre"])
		fmt.Println(strings.Repeat(" ", int(indent)) + " year : " + item["year"])
	}
}
