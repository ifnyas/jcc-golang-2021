package main

import (
	"fmt"
	"strconv"
)

func main() {
	// soal 1
	fmt.Println("-----SOAL 1-----")
	for i := 1; i <= 20; i++ {
		if i%2 == 1 && i%3 == 0 {
			fmt.Println(i, "- I Love Coding")
		} else if i%2 == 0 {
			fmt.Println(i, "- Candradimuka")
		} else {
			fmt.Println(i, "- JCC")
		}
	}

	// soal 2
	fmt.Println("-----SOAL 2-----")
	for i := 1; i <= 7; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	// soal 3
	fmt.Println("-----SOAL 3-----")
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	newKalimat := kalimat[2:]

	fmt.Println(newKalimat)

	// soal 4
	fmt.Println("-----SOAL 4-----")
	var sayuran = []string{}

	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	for idx, sayur := range sayuran {
		fmt.Println(strconv.Itoa(idx) + ". " + sayur)
	}

	// soal 5
	fmt.Println("-----SOAL 5-----")
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	var volumeBalok int = 1

	for key, item := range satuan {
		fmt.Println(key, "=", item)
		volumeBalok *= item
	}

	fmt.Println("Volume Balok =", volumeBalok)
}
