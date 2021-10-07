package main

import (
	"fmt"
	"strconv"
)

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

type description interface {
	showDescription() string
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return float64((2 * (b.panjang + b.lebar)) + (2 * (b.panjang + b.tinggi)) + (2 * (b.tinggi + b.lebar)))
}

func (t tabung) volume() float64 {
	if int(t.jariJari)%7 == 0 {
		return float64(22 / 7 * t.jariJari * t.jariJari * t.tinggi)
	} else {
		return float64(3.14 * t.jariJari * t.jariJari * t.tinggi)
	}
}

func (t tabung) luasPermukaan() float64 {
	if int(t.jariJari)%7 == 0 {
		return float64(2 * 22 / 7 * t.jariJari * (t.jariJari + t.tinggi))
	} else {
		return float64(2 * 3.14 * t.jariJari * (t.jariJari + t.tinggi))
	}
}

func (p phone) showDescription() (desc string) {
	var colors string

	for index, item := range p.colors {
		if index == 0 {
			colors += item
		} else {
			colors += ", " + item
		}
	}

	desc = "name : " + p.name + "\n" +
		"brand : " + p.brand + "\n" +
		"year : " + strconv.Itoa(p.year) + "\n" +
		"color: " + colors
	return
}

func luasPersegi(sisi uint, tampilkan bool) interface{} {
	switch {
	case sisi == 0 && tampilkan == false:
		return nil
	case sisi == 0 && tampilkan == true:
		return "Maaf anda belum menginput sisi dari persegi"
	case tampilkan == false:
		return sisi * sisi
	default:
		return "luas persegi dengan sisi " + strconv.Itoa(int(sisi)) + "cm adalah " + strconv.Itoa(int(sisi*sisi)) + " cm"
	}
}

func main() {

	// soal 1
	fmt.Println("-----SOAL 1-----")

	// bangun datar

	var hitungSegitigaPertama hitungBangunDatar = segitigaSamaSisi{alas: 6, tinggi: 10}

	fmt.Println(hitungSegitigaPertama.luas())
	fmt.Println(hitungSegitigaPertama.keliling())

	var hitungPersegiPanjangPertama hitungBangunDatar = persegiPanjang{lebar: 6, panjang: 10}

	fmt.Println(hitungPersegiPanjangPertama.luas())
	fmt.Println(hitungPersegiPanjangPertama.keliling())

	// bangun ruang
	var hitungBalok hitungBangunRuang = balok{lebar: 6, panjang: 10, tinggi: 5}
	fmt.Println(hitungBalok.volume())
	fmt.Println(hitungBalok.luasPermukaan())

	var hitungTabung hitungBangunRuang = tabung{jariJari: 7, tinggi: 5}
	fmt.Println(hitungTabung.volume())
	fmt.Println(hitungTabung.luasPermukaan())

	// soal 2
	fmt.Println("-----SOAL 2-----")

	var hape description = phone{name: "Samsung Galaxy Note 20", brand: "Samsung", year: 2020, colors: []string{"Bronze", "White", "Black"}}

	fmt.Println(hape.showDescription())

	// soal 3
	fmt.Println("-----SOAL 3-----")

	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))

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
