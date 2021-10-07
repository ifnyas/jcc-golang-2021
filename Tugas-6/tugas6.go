package main

import (
	"fmt"
)

/* ----- soal 1-----*/

func perbaruiHasilLingkaran(luas, keliling *float64, jariJari int) {
	if jariJari%7 == 0 {
		*luas = float64(22 / 7 * jariJari * jariJari)
		*keliling = float64(2 * 22 / 7 * jariJari)
	} else {
		*luas = 3.14 * float64(jariJari*jariJari)
		*keliling = 3.14 * float64(2*jariJari)
	}
}

/* ----- soal 2-----*/

func introduce(sentence *string, name, gender, job, age string) {
	*sentence = "Pak " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
}

/* ----- soal 3-----*/

func tambahBuah(buah *[]string, namaBuah string) {
	*buah = append(*buah, namaBuah)
}

/* ----- soal 4-----*/

func tambahDataFilm(title, duration, genre, year string, dataFilm *[]map[string]string) {
	data := map[string]string{
		"title": title,
		"jam":   duration,
		"tahun": year,
		"genre": genre,
	}

	*dataFilm = append(*dataFilm, data)
}
func main() {
	// soal 1
	fmt.Println("-----SOAL 1-----")
	var luasLigkaran float64
	var kelilingLingkaran float64

	perbaruiHasilLingkaran(&luasLigkaran, &kelilingLingkaran, 7)

	fmt.Println("Luas Lingkaran : ", luasLigkaran)
	fmt.Println("Keliling Lingkaran :", kelilingLingkaran)

	perbaruiHasilLingkaran(&luasLigkaran, &kelilingLingkaran, 10)

	fmt.Println("Luas Lingkaran : ", luasLigkaran)
	fmt.Println("Keliling Lingkaran :", kelilingLingkaran)

	// soal 2
	fmt.Println("-----SOAL 2-----")

	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
	// soal 3
	fmt.Println("-----SOAL 3-----")

	var buah = []string{}

	tambahBuah(&buah, "Jeruk")
	tambahBuah(&buah, "Semangka")
	tambahBuah(&buah, "Mangga")
	tambahBuah(&buah, "Strawberry")
	tambahBuah(&buah, "Durian")
	tambahBuah(&buah, "Manggis")
	tambahBuah(&buah, "Alpukat")

	for index, item := range buah {
		fmt.Print(index + 1)
		fmt.Println(". " + item)
	}
	// soal 4
	fmt.Println("-----SOAL 4-----")

	var dataFilm = []map[string]string{}
	// buatlah closure function disini

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for index, item := range dataFilm {
		fmt.Print(index + 1)
		fmt.Print(". ")
		i := 1
		for key, value := range item {
			if i == 1 {
				fmt.Println(key + " : " + value)
			} else {
				fmt.Println("   " + key + " : " + value)
			}
			i++
		}
		fmt.Println()
	}

}
