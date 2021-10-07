package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"
)

func printPhone(number int, name string, wg *sync.WaitGroup) {
	fmt.Println(strconv.Itoa(number) + ". " + name)
	time.Sleep(time.Second)
	wg.Done()
}

func getMovies(movieCh chan string, movie ...string) {
	movieCh <- "List Movie:"
	for index, item := range movie {
		movieCh <- strconv.Itoa(index+1) + ". " + item
	}
	close(movieCh)
}

func luasLingkaran(jariJari float64, ch chan float64) {
	ch <- math.Round(math.Pi * jariJari * jariJari)
}

func kelilingLingkaran(jariJari float64, ch chan float64) {
	ch <- math.Round(2 * math.Pi * jariJari)
}

func volumeTabung(jariJari float64, tinggi float64, ch chan float64) {
	ch <- math.Round(math.Pi * jariJari * jariJari * tinggi)
}

func luasPersegiPanjang(p int, l int, ch chan int) {
	ch <- p * l
}

func kelilingPersegiPanjang(p int, l int, ch chan int) {
	ch <- 2 * (p + l)

}

func volumeBalok(p int, l int, t int, ch chan int) {
	ch <- p * l * t
}

func main() {
	// soal 1
	fmt.Println("-----SOAL 1------")
	var wg sync.WaitGroup
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

	for index, phone := range phones {
		wg.Add(1)
		go printPhone(index+1, phone, &wg)
		wg.Wait()
	}

	// soal 2
	fmt.Println("-----SOAL 2------")

	var movies = []string{
		"Harry Potter",
		"LOTR",
		"SpiderMan",
		"Logan",
		"Avengers",
		"Insidious",
		"Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}

	// soal 2
	fmt.Println("-----SOAL 3------")

	hitungLuas := make(chan float64)
	hitungKeliling := make(chan float64)
	hitungVolume := make(chan float64)

	sliceOfJariJari := []float64{8, 14, 20}

	for _, item := range sliceOfJariJari {

		go luasLingkaran(item, hitungLuas)
		hasil := <-hitungLuas
		fmt.Println("luas lingkaran dari jari-jari", item, "adalah", hasil)

		go kelilingLingkaran(item, hitungKeliling)
		hasil2 := <-hitungKeliling
		fmt.Println("keliling lingkaran dari jari-jari", item, "adalah", hasil2)

		go volumeTabung(item, 10, hitungVolume)
		hasil3 := <-hitungVolume
		fmt.Println("volume balok dari jari-jari", item, "dan tinggi 10 adalah", hasil3)
	}

	// soal 4
	fmt.Println("-----SOAL 4------")

	hitungLuasPP := make(chan int)
	go luasPersegiPanjang(7, 5, hitungLuasPP)

	hitungKelilingPP := make(chan int)
	go kelilingPersegiPanjang(7, 5, hitungKelilingPP)

	hitungVolumeBalok := make(chan int)
	go volumeBalok(7, 5, 10, hitungVolumeBalok)

	for i := 1; i <= 3; i++ {
		select {
		case luas := <-hitungLuasPP:
			fmt.Println("Panjang = 7, Lebar = 5, Luas =", luas)
		case keliling := <-hitungKelilingPP:
			fmt.Println("Panjang = 7, Lebar = 5, Keliling =", keliling)
		case volume := <-hitungVolumeBalok:
			fmt.Println("Panjang = 7, Lebar = 5, tinggi=10, Volume Balok =", volume)
		}
	}

}
