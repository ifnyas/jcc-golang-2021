package main

import "fmt"

type buah struct {
	nama, warna string
	adaBijinya  bool
	harga       int
}

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type movie struct {
	title, genre   string
	year, duration int
}

func (s segitiga) luasSegitiga() int {
	return s.alas * s.tinggi / 2
}

func (p persegi) luasPersegi() int {
	return p.sisi * p.sisi
}

func (pp persegiPanjang) luasPersegiPanjang() int {
	return pp.panjang * pp.lebar
}

func (p *phone) addColor(color string) {
	p.colors = append(p.colors, color)
}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	data := movie{
		title:    title,
		duration: duration,
		year:     year,
		genre:    genre,
	}

	*dataFilm = append(*dataFilm, data)
}

func main() {
	// soal 1
	buah1 := buah{nama: "Nanas", warna: "Kuning", adaBijinya: false, harga: 9000}
	buah2 := buah{nama: "Jeruk", warna: "Oranye", adaBijinya: true, harga: 8000}
	buah3 := buah{nama: "Semangka", warna: "Hijau  & Merah", adaBijinya: true, harga: 10000}
	buah4 := buah{nama: "Pisang", warna: "Kuning", adaBijinya: false, harga: 5000}

	fmt.Println(buah1)
	fmt.Println(buah2)
	fmt.Println(buah3)
	fmt.Println(buah4)

	// soal 2
	segitiga1 := segitiga{}
	segitiga1.alas = 8
	segitiga1.tinggi = 10

	fmt.Println(segitiga1.luasSegitiga())

	persegi1 := persegi{}
	persegi1.sisi = 8
	fmt.Println(persegi1.luasPersegi())

	persegiPanjang1 := persegiPanjang{}
	persegiPanjang1.panjang = 8
	persegiPanjang1.lebar = 5
	fmt.Println(persegiPanjang1.luasPersegiPanjang())

	// soal 3
	samsung := phone{name: "Samsung Galaxy Note 20", brand: "Samsung", year: 2020}

	fmt.Println(samsung)

	samsung.addColor("Silver")
	samsung.addColor("Bronze")
	samsung.addColor("Gold")

	fmt.Println(samsung)

	// soal 4
	fmt.Println("-----SOAL 4-----")

	var dataFilm = []movie{}
	// buatlah closure function disini

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for index, item := range dataFilm {
		fmt.Print(index + 1)
		fmt.Println(". title:", item.title)
		fmt.Println("   duration:", item.duration/60, "jam")
		fmt.Println("   genre:", item.genre)
		fmt.Println("   year:", item.year)
		fmt.Println()
	}

}
