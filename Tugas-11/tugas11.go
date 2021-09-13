package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"
)

func main() {
	// soal 1
	phones := []string{}
	addPhone := func(s string, p *[]string) {
		phones = append(*p, s)
	}

	phonesToAdd := []string{
		"Xiaomi",
		"Asus",
		"IPhone",
		"Samsung",
		"Oppo",
		"Realme",
		"Vivo"}
	sort.Strings(phonesToAdd)

	for i, p := range phonesToAdd {
		addPhone(p, &phones)
		result := strconv.Itoa(i+1) + ". " + phones[i]
		fmt.Println(result)
		time.Sleep(1 * time.Second)
	}

	// soal 2
	kelilingO := func(r int) float64 {
		return math.Round(2 * math.Pi * float64(r))
	}
	luasO := func(r int) float64 {
		return math.Round(math.Pi * math.Pow(float64(r), 2))
	}
	jariJariO := []int{7, 10, 15}
	for _, n := range jariJariO {
		fmt.Println("r:", n, "| Luas:", luasO(n))
		fmt.Println("r:", n, "| Keliling:", kelilingO(n))
	}

	// soal 3
	kelilingPp := func(p int, l int) int {
		return 2 * (p + l)
	}
	luasPp := func(p int, l int) int {
		return p * l
	}

	p := flag.Int("panjang", 2, "Masukkan Panjang")
	l := flag.Int("lebar", 3, "Masukkan Lebar")
	flag.Parse()

	fmt.Println("p:", *p, ", l:", *l, "| Luas:", luasPp(*p, *l))
	fmt.Println("p:", *p, ", l:", *l, "| Keliling:", kelilingPp(*p, *l))
}
