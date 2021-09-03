package main

import (
	"fmt"
	"strings"
)

func main() {
	// soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	// soal 2
	introduce := func(name string, gender string, job string, age string) string {
		title := ""
		if gender == "laki-laki" {
			title = "Pak"
		} else {
			title = "Bu"
		}
		return title + " " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	}

	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)

	// soal 3
	buahFavorit := func(name string, fruits ...string) string {
		res := "halo nama saya " + strings.ToLower(name) + " dan buah favorit saya adalah "
		for _, fruit := range fruits {
			res += "\"" + fruit + "\", "
		}
		return strings.Trim(res, ", ")
	}
	buah := []string{"semangka", "jeruk", "melon", "pepaya"}
	buahFavoritJohn := buahFavorit("John", buah...)
	fmt.Println(buahFavoritJohn)

	// soal 4
	var dataFilm = []map[string]string{}
	tambahDataFilm := func(str ...string) {
		film := map[string]string{
			"title": str[0],
			"jam":   strings.Trim(str[1], " jam"),
			"genre": str[2],
			"tahun": str[3],
		}
		dataFilm = append(dataFilm, film)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}

func luasPersegiPanjang(p int, l int) int {
	return p * l
}

func kelilingPersegiPanjang(p int, l int) int {
	return (p + l) * 2
}

func volumeBalok(p int, l int, t int) int {
	return p * l * t
}
