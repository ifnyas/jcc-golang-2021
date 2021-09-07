package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	type buah struct {
		nama, warna string
		adaBijinya  bool
		harga       int
	}
	buahan := [...]buah{
		{"Nanas", "Kuning", false, 9000},
		{"Jeruk", "Oranye", true, 8000},
		{"Semangka", "Hijau & Merah", true, 10000},
		{"Pisang", "Kuning", false, 5000},
	}
	for _, item := range buahan {
		fmt.Println(item)
	}

	// soal 2
	iSegitiga := segitiga{2, 3}
	iPersegi := persegi{2}
	iPersegiPanjang := persegiPanjang{2, 3}

	fmt.Println(iSegitiga.luas())
	fmt.Println(iPersegi.luas())
	fmt.Println(iPersegiPanjang.luas())

	// soal 3
	iPhone := phone{"My Phone", "iBrand", 2021, []string{}}
	iPhone.addColor("Red")
	fmt.Println(iPhone)

	// soal 4
	var dataFilm = []movie{}
	tambahDataFilm := func(title string, duration int, genre string, year int, df *[]movie) {
		data := movie{title, genre, duration, year}
		*df = append(*df, data)
	}
	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, item := range dataFilm {
		indent := math.Floor(float64(i)/10) + 2
		durationStr := strconv.Itoa(item.duration/60) + " jam"
		fmt.Println(strconv.Itoa(i+1) + ". title : " + item.title)
		fmt.Println(strings.Repeat(" ", int(indent)) + " duration : " + durationStr)
		fmt.Println(strings.Repeat(" ", int(indent)) + " genre : " + item.genre)
		fmt.Println(strings.Repeat(" ", int(indent)) + " year : " + strconv.Itoa(item.year))
	}
}

// soal 2
type segitiga struct {
	alas, tinggi int
}
type persegi struct {
	sisi int
}
type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) luas() int {
	return s.alas * s.tinggi / 2
}
func (s persegi) luas() int {
	return s.sisi * s.sisi
}
func (s persegiPanjang) luas() int {
	return s.panjang * s.lebar
}

// soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) addColor(color string) {
	p.colors = append(p.colors, color)
}

// soal 4
type movie struct {
	title, genre   string
	duration, year int
}
