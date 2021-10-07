package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	kata1 := "Jabar"
	kata2 := "Coding"
	kata3 := "Camp"
	kata4 := "Golang"
	kata5 := "2021"
	fmt.Println("-----SOAL 1-----")
	fmt.Println(kata1, kata2, kata3, kata4, kata5)

	// soal 2
	halo := "Halo Dunia"
	fmt.Println("-----SOAL 2-----")
	fmt.Println(strings.Replace(halo, "Dunia", "Golang", 1))

	// soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	kataKetiga = strings.Replace(kataKetiga, string(kataKetiga[len(kataKetiga)-1]), "", 1) + strings.ToUpper(string(kataKetiga[len(kataKetiga)-1]))

	fmt.Println("-----SOAL 3-----")
	fmt.Println(kataPertama, strings.Title(kataKedua), kataKetiga, strings.ToUpper(kataKeempat))

	// soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	angka1, _ := strconv.Atoi(angkaPertama)
	angka2, _ := strconv.Atoi(angkaKedua)
	angka3, _ := strconv.Atoi(angkaKetiga)
	angka4, _ := strconv.Atoi(angkaKeempat)

	fmt.Println("-----SOAL 4-----")
	fmt.Println(angka1 + angka2 + angka3 + angka4)

	// soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	kalimatModif := `"` + strings.Replace(kalimat, "halo", "hi", -1) + `"`

	fmt.Println("-----SOAL 5-----")
	fmt.Println(kalimatModif, angka)
}
