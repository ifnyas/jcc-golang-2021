package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Soal dimulai dari nomor 2 dalam materi tugas

	// soal 1(+1)
	for i := 1; i <= 20; i++ {
		str := ""
		switch {
		case i%2 == 1:
			if i%3 == 0 {
				str = "I Love Coding"
			} else {
				str = "JCC"
			}
		default:
			str = "Candradimuka"
		}
		fmt.Println(i, "-", str)
	}

	// soal 2(+1)
	for i := 1; i <= 7; i++ {
		fmt.Println(strings.Repeat("#", i))
	}

	// soal 3(+1)
	kalimat := [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2:])

	// soal 4(+1)
	sayuran := []string{}
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")
	for i, sayur := range sayuran {
		fmt.Println(strconv.Itoa(i+1) + ". " + sayur)
	}

	// soal 5(+1)
	satuan := map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	satuan["luas"] = satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
	for satu := range satuan {
		fmt.Println(satu, "=", satuan[satu])
	}
}
