package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

func tambahPhone(phones *[]string, name string) {
	*phones = append(*phones, name)
}

func luasLingkaran(jariJari float64) float64 {
	return math.Round(math.Pi * jariJari * jariJari)
}

func kelilingLingkaran(jariJari float64) float64 {
	return math.Round(2 * math.Pi * jariJari)
}

func main() {
	// soal 1
	fmt.Println("-----SOAL 1------")
	var phones = []string{}
	tambahPhone(&phones, "Xiaomi")
	tambahPhone(&phones, "Asus")
	tambahPhone(&phones, "IPhone")
	tambahPhone(&phones, "Samsung")
	tambahPhone(&phones, "Oppo")
	tambahPhone(&phones, "Realme")
	tambahPhone(&phones, "Vivo")

	for index, phone := range phones {
		fmt.Print(index + 1)
		fmt.Println(".", phone)
		time.Sleep(time.Second)
	}

	// soal 2
	fmt.Println("-----SOAL 2------")

	fmt.Println("Luas Lingkaran dari 7 adalah", luasLingkaran(7))
	fmt.Println("Luas Lingkaran dari 10 adalah", luasLingkaran(10))
	fmt.Println("Luas Lingkaran dari 15 adalah", luasLingkaran(15))

	fmt.Println("Keliling Lingkarang dari 7 adalah", kelilingLingkaran(7))
	fmt.Println("Keliling Lingkarang dari 10 adalah", kelilingLingkaran(10))
	fmt.Println("Keliling Lingkarang dari 15 adalah", kelilingLingkaran(15))

	// soal 3
	fmt.Println("-----SOAL 2------")
	var panjang = flag.Int("panjang", 0, "tulis panjang persegi Panjang")
	var lebar = flag.Int("lebar", 0, "tulis lebar persegi Panjang")

	flag.Parse()
	fmt.Println("Panjang = ", *panjang)
	fmt.Println("Lebar = ", *lebar)
	fmt.Println("Luas = ", (*panjang)*(*lebar))
	fmt.Println("Keliling = ", 2*(*panjang+*lebar))

}
