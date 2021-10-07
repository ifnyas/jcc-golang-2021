package main

import (
	"errors"
	"fmt"
	"strconv"
)

func endApp() {
	fmt.Println("----defer soal 2-----")
	message := recover()
	fmt.Println(message)
}

func showPanic() {
	defer endApp()
	panic("Terjadi ERROR")
}

func kelilingSegitigaSamaSisi(angka int, kalimat bool) (string, error) {
	switch {
	case angka == 0 && kalimat == false:
		{
			showPanic()
			return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}
	case angka == 0 && kalimat == true:
		{
			return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}
	case kalimat == false:
		{
			return strconv.Itoa(angka), nil
		}
	default:
		{
			return "keliling segitiga sama sisinya dengan sisi " + strconv.Itoa(angka) + " cm adalah " + strconv.Itoa(angka*3) + " cm", nil

		}
	}
}

func showKalimat(kalimat string, tahun int) {
	fmt.Println("----defer soal 1-----")
	fmt.Println(kalimat, tahun)
}

func tambahAngka(tambahanAngka int, angka *int) {
	*angka += tambahanAngka
}

func totalAngka(angka *int) {
	fmt.Println("----defer soal 3-----")
	fmt.Println(*angka)
}

func main() {
	// deklarasi variabel angka ini simpan di baris pertama func main
	angka := 1
	// soal 1
	defer showKalimat("Candradimuka Jabar Coding Camp", 2021)
	defer totalAngka(&angka)

	// soal 2
	keliling1, err := kelilingSegitigaSamaSisi(4, true)

	if err == nil {
		fmt.Println(keliling1)
	} else {
		fmt.Println(err)
	}

	keliling2, err := kelilingSegitigaSamaSisi(8, false)

	if err == nil {
		fmt.Println(keliling2)
	} else {
		fmt.Println(err)
	}

	keliling3, err := kelilingSegitigaSamaSisi(0, true)

	if err == nil {
		fmt.Println(keliling3)
	} else {
		fmt.Println(err)
	}

	keliling4, err := kelilingSegitigaSamaSisi(0, false)

	if err == nil {
		fmt.Println(keliling4)
	} else {
		fmt.Println(err)
	}

	tambahAngka(7, &angka)

	tambahAngka(6, &angka)

	tambahAngka(-1, &angka)

	tambahAngka(9, &angka)

}
