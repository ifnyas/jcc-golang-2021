package main

import (
	"fmt"
)

/* -----soal 1-----*/

func luasPersegiPanjang(p, l int) int {
	return p * l
}

func kelilingPersegiPanjang(p, l int) int {
	return 2 * (p + l)
}

func volumeBalok(p, l, t int) int {
	return p * l * t
}

/* -----soal 1-----*/

/* -----soal 2-----*/
func introduce(name, gender, job, age string) (sentence string) {
	prefix := "Pak"

	if gender == "perempuan" {
		prefix = "Bu"
	}
	sentence = prefix + " " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	return
}

/* -----soal 2-----*/

func buahFavorit(name string, fruits ...string) string {

	sentence := "halo nama saya " + name + " dan buah favorit saya adalah"

	for index, fruit := range fruits {
		if index == 0 {
			sentence += ` "` + fruit + `"`
		} else {
			sentence += `, "` + fruit + `"`
		}
	}

	return sentence
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

}

func main() {
	// soal 1
	fmt.Println("-----SOAL 1-----")

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
	fmt.Println("-----SOAL 2-----")
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// soal 3
	fmt.Println("-----SOAL 3-----")
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	// soal 4
	fmt.Println("-----SOAL 4-----")

	var dataFilm = []map[string]string{}
	// buatlah closure function disini

	tambahDataFilm := func(title, duration, genre, year string) {
		data := map[string]string{
			"title": title,
			"jam":   duration,
			"tahun": year,
			"genre": genre,
		}

		dataFilm = append(dataFilm, data)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

}
