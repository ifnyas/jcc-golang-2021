package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	str1 := "Jabar"
	str2 := "Coding"
	str3 := "Camp"
	str4 := "Golang"
	str5 := "2021"
	fmt.Println(str1, str2, str3, str4, str5)

	// soal 2
	halo := "Halo Dunia"
	haloReplaced := strings.Replace(halo, "Dunia", "Golang", -1)
	fmt.Println(haloReplaced)

	// soal 3
	kataPertama := "saya"
	kataKedua := "senang"
	kataKetiga := "belajar"
	kataKeempat := "golang"

	kataKeduaModded := strings.Title(kataKedua)
	kataKetigaLast := kataKetiga[len(kataKetiga)-1:]
	kataKetigaModded := strings.Replace(
		kataKetiga,
		kataKetigaLast,
		strings.ToUpper(kataKetigaLast),
		strings.LastIndex(kataKetiga, kataKetigaLast))
	kataKeempatModded := strings.ToUpper(kataKeempat)

	kataArr := []string{
		kataPertama,
		kataKeduaModded,
		kataKetigaModded,
		kataKeempatModded}
	kataJoined := strings.Join(kataArr, " ")
	fmt.Println(kataJoined)

	// soal 4
	angkaPertama := "8"
	angkaKedua := "5"
	angkaKetiga := "6"
	angkaKeempat := "7"

	angkaPertamaInt, err := strconv.Atoi(angkaPertama)
	if err != nil {
		panic(err)
	}

	angkaKeduaInt, err := strconv.Atoi(angkaKedua)
	if err != nil {
		panic(err)
	}

	angkaKetigaInt, err := strconv.Atoi(angkaKetiga)
	if err != nil {
		panic(err)
	}
	angkaKeempatInt, err := strconv.Atoi(angkaKeempat)
	if err != nil {
		panic(err)
	}

	angkaSum := angkaPertamaInt + angkaKeduaInt + angkaKetigaInt + angkaKeempatInt
	fmt.Println(angkaSum)

	// soal 5
	kalimat := "halo halo bandung"
	angka := 2021
	output := "\"" + strings.Replace(kalimat, "halo", "Hi", -1) + "\" - " + strconv.Itoa(angka)
	fmt.Println(output)
}
