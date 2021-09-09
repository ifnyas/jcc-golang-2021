package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// soal 1
	runSoalSatu()

	// soal 2
	runSoalDua()

	// soal 3
	runSoalTiga()
}

// soal 1
func runSoalSatu() {
	defer deferSoalSatu("Candradimuka Jabar Coding Camp", 2021)
}

func deferSoalSatu(s string, i int) {
	fmt.Println(s, strconv.Itoa(i))
}

// soal 2
func runSoalDua() {
	defer soalDuaPanicLogging()

	kelilingSegitigaSamaSisi := func(n int, b bool) (string, error) {
		var result string
		var err error

		switch {
		case n > 0 && b:
			result = "keliling segitiga sama sisinya dengan sisi " +
				strconv.Itoa(n) + " cm adalah " +
				strconv.Itoa(n*3) + " cm"
		case n > 0 && !b:
			result = strconv.Itoa(n * 3)
		case n == 0 && b:
			err = errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		case n == 0 && !b:
			panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}

		soalDuaLogging(result, err) // why logging here if you return the values?
		return result, err
	}

	// I want to type these funcs exactly like on the learning material
	kelilingSegitigaSamaSisi(4, true)
	kelilingSegitigaSamaSisi(8, false)
	kelilingSegitigaSamaSisi(0, true)
	kelilingSegitigaSamaSisi(0, false)
	// but we can always put it like soalDuaLogging(kelilingSegitigaSamaSisi(4, true))
}

func soalDuaLogging(s string, e error) {
	var r interface{}
	switch {
	case s == "" && e == nil:
		r = recover()
	case e != nil:
		r = e.Error()
	default:
		r = s
	}
	fmt.Println(r)
}

func soalDuaPanicLogging() {
	fmt.Println(recover())
}

func runSoalTiga() {
	angka := 1

	defer soalTigaLogging(&angka)

	tambahAngka := func(n int, angka *int) {
		*angka += n
	}

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)
}

func soalTigaLogging(n *int) {
	fmt.Println(strconv.Itoa(*n))
}
