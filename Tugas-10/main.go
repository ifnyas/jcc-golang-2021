package main

import (
	. "Tugas-10/soal1"
	. "Tugas-10/soal2"
	"Tugas-10/soal3"
	"fmt"
	"strconv"
)

func main() {

	// soal 1
	fmt.Println("-----SOAL 1-----")

	// bangun datar

	var hitungSegitigaPertama HitungBangunDatar = SegitigaSamaSisi{Alas: 6, Tinggi: 10}

	fmt.Println(hitungSegitigaPertama.Luas())
	fmt.Println(hitungSegitigaPertama.Keliling())

	var hitungPersegiPanjangPertama HitungBangunDatar = PersegiPanjang{Lebar: 6, Panjang: 10}

	fmt.Println(hitungPersegiPanjangPertama.Luas())
	fmt.Println(hitungPersegiPanjangPertama.Keliling())

	// bangun ruang
	var hitungBalok HitungBangunRuang = Balok{Lebar: 6, Panjang: 10, Tinggi: 5}
	fmt.Println(hitungBalok.Volume())
	fmt.Println(hitungBalok.LuasPermukaan())

	var hitungTabung HitungBangunRuang = Tabung{JariJari: 7, Tinggi: 5}
	fmt.Println(hitungTabung.Volume())
	fmt.Println(hitungTabung.LuasPermukaan())

	// soal 2
	fmt.Println("-----SOAL 2-----")

	var hape Description = Phone{Name: "Samsung Galaxy Note 20", Brand: "Samsung", Year: 2020, Colors: []string{"Bronze", "White", "Black"}}

	fmt.Println(hape.ShowDescription())

	// soal 3
	fmt.Println("-----SOAL 3-----")

	fmt.Println(soal3.LuasPersegi(4, true))

	fmt.Println(soal3.LuasPersegi(8, false))

	fmt.Println(soal3.LuasPersegi(0, true))

	fmt.Println(soal3.LuasPersegi(0, false))

	// soal 4

	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	kalimat := "" + prefix.(string)

	kumpulanAngka := []int{}

	kumpulanAngka = append(kumpulanAngka, kumpulanAngkaPertama.([]int)[0], kumpulanAngkaPertama.([]int)[1])
	kumpulanAngka = append(kumpulanAngka, kumpulanAngkaKedua.([]int)[0], kumpulanAngkaKedua.([]int)[1])

	jumlah := 0

	for index, item := range kumpulanAngka {
		if index == 0 {
			kalimat += strconv.Itoa(item)
		} else {
			kalimat += "+" + strconv.Itoa(item)
		}
		jumlah += item
	}

	kalimat += "=" + strconv.Itoa(jumlah)

	fmt.Println(kalimat)

}
