package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	var iDatar hitungBangunDatar
	var iRuang hitungBangunRuang

	iDatar = segitigaSamaSisi{3, 3}
	fmt.Println("Luas S3:", iDatar.luas())
	fmt.Println("Keliling S3:", iDatar.keliling())

	iDatar = persegiPanjang{2, 3}
	fmt.Println("Luas PP:", iDatar.luas())
	fmt.Println("Keliling PP:", iDatar.keliling())

	iRuang = balok{2, 3, 4}
	fmt.Println("Volume Balok:", iRuang.volume())
	fmt.Println("Luas P. Balok:", iRuang.luasPermukaan())

	iRuang = tabung{2, 3}
	fmt.Println("Volume Tabung:", fmt.Sprintf("%.2f", iRuang.volume()))
	fmt.Println("Luas P. Tabung:", fmt.Sprintf("%.2f", iRuang.luasPermukaan()))

	// soal 2
	var iPhone = phone{
		"Galaxy Note S20",
		"Samsung",
		2020,
		[]string{"Mystic Bronze", "Mystic White", "Mystic Black"}}
	fmt.Println(iPhone.getData())

	// soal 3
	var result interface{}
	var luasPersegi = func(n int, b bool) interface{} {
		switch {
		case n > 0 && b:
			result = "luas persegi dengan sisi " +
				strconv.Itoa(n) + " cm adalah " +
				strconv.Itoa(n*n) + " cm"
		case n > 0 && !b:
			result = n * n
		case n == 0 && b:
			result = "Maaf anda belum menginput sisi dari persegi"
		case n == 0 && !b:
			result = nil
		}
		return result
	}
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	// soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	var kumpulanAngka interface{} = []int{}
	var total int
	for _, n := range kumpulanAngkaPertama.([]int) {
		kumpulanAngka = append(kumpulanAngka.([]int), n)
		total += n
	}
	for _, n := range kumpulanAngkaKedua.([]int) {
		kumpulanAngka = append(kumpulanAngka.([]int), n)
		total += n
	}
	kumpulanAngka = strings.Trim(
		strings.Join(
			strings.Split(
				fmt.Sprint(
					kumpulanAngka.([]int)),
				" "),
			" + "),
		"[]")
	var kumpulanStr = prefix.(string) + kumpulanAngka.(string) + " = " + strconv.Itoa(total)
	fmt.Println(kumpulanStr)
}

// soal 1
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

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

func (s persegiPanjang) luas() int {
	return s.panjang * s.lebar
}

func (s persegiPanjang) keliling() int {
	return 2 * (s.panjang + s.lebar)
}

func (s balok) volume() float64 {
	return float64(s.lebar) * float64(s.panjang) * float64(s.tinggi)
}

func (s balok) luasPermukaan() float64 {
	return (2 * (float64(s.lebar) + float64(s.panjang))) +
		(2 * (float64(s.lebar) + float64(s.tinggi))) +
		(2 * (float64(s.tinggi) + float64(s.panjang)))
}

func (s tabung) volume() float64 {
	return math.Pi * math.Pow(s.jariJari, 2) * s.tinggi
}

func (s tabung) luasPermukaan() float64 {
	return 2 * math.Pi * s.jariJari * (s.jariJari + s.tinggi)
}

// soal 2
type phone struct {
	name, brand string
	year        int
	colors      []string
}

var iFace interface {
	getData() string
	printData()
}

func (p phone) getData() string {
	var result string
	result += "name : " + p.name
	result += "\nbrand : " + p.brand
	result += "\nyear : " + strconv.Itoa(p.year)
	result += "\ncolors : " + strings.Join(p.colors, ", ")
	return result
}
