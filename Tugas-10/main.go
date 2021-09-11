package main

import (
	"fmt"
	"strconv"
	"strings"

	. "jcc-golang-2021/Tugas-10/library"
)

func main() {
	var iDatar HitungBangunDatar
	var iRuang HitungBangunRuang

	iDatar = SegitigaSamaSisi{3, 3}
	fmt.Println("Luas S3:", iDatar.Luas())
	fmt.Println("Keliling S3:", iDatar.Keliling())

	iDatar = PersegiPanjang{2, 3}
	fmt.Println("Luas PP:", iDatar.Luas())
	fmt.Println("Keliling PP:", iDatar.Keliling())

	iRuang = Balok{2, 3, 4}
	fmt.Println("Volume Balok:", iRuang.Volume())
	fmt.Println("Luas P. Balok:", iRuang.LuasPermukaan())

	iRuang = Tabung{2, 3}
	fmt.Println("Volume Tabung:", fmt.Sprintf("%.2f", iRuang.Volume()))
	fmt.Println("Luas P. Tabung:", fmt.Sprintf("%.2f", iRuang.LuasPermukaan()))

	// soal 2
	var iPhone = Phone{
		"Galaxy Note S20",
		"Samsung",
		2020,
		[]string{"Mystic Bronze", "Mystic White", "Mystic Black"}}
	fmt.Println(iPhone.GetData())

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
