package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"sync"
)

func main() {
	// soal 1
	var wg sync.WaitGroup

	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	sort.Strings(phones)

	var printPhone = func(result string) {
		//time.Sleep(1 * time.Second)
		fmt.Println(result)
		wg.Done()
	}

	for i, p := range phones {
		result := strconv.Itoa(i+1) + ". " + p
		wg.Add(1)
		go printPhone(result)
		wg.Wait()
	}

	// soal 2
	var movies = []string{
		"Harry Potter",
		"LOTR",
		"SpiderMan",
		"Logan",
		"Avengers",
		"Insidious",
		"Toy Story"}

	var moviesChannel = make(chan string)

	var getMovies = func(ch chan string, movies ...string) {
		for i, movie := range movies {
			var result = strconv.Itoa(i+1) + ". " + movie
			ch <- result
		}
		close(ch)
	}

	go getMovies(moviesChannel, movies...)

	fmt.Println("List Movies :")
	for value := range moviesChannel {
		fmt.Println(value)
	}

	// soal 3
	var tinggiTabung = 10
	var jariJariO = []int{8, 14, 20}

	var luasO = func(r int) float64 {
		return math.Round(math.Pi * math.Pow(float64(r), 2))
	}
	var kelilingO = func(r int) float64 {
		return math.Round(2 * math.Pi * float64(r))
	}
	var volumeTabung = func(r int) float64 {
		return math.Round(math.Pi * math.Pow(float64(r), 2) * float64(tinggiTabung))
	}
	var inputChan = func(ch chan int, rs ...int) {
		for _, r := range rs {
			ch <- r
		}
		close(ch)
	}

	var jariJariChannel = make(chan int)
	go inputChan(jariJariChannel, jariJariO...)

	for r := range jariJariChannel {
		fmt.Println("r:", r, "| Luas Lingkaran:", luasO(r))
		fmt.Println("r:", r, "| Keliling Lingkaran:", kelilingO(r))
		fmt.Println("r:", r, "| Volume Tabung:", volumeTabung(r))
	}

	// soal 4
	var kelilingPp = func(p int, l int, ch chan int) {
		ch <- 2 * (p + l)
	}
	var luasPp = func(p int, l int, ch chan int) {
		ch <- p * l
	}
	var volumeBalok = func(p int, l int, t int, ch chan int) {
		ch <- p * l * t
	}

	var kChan = make(chan int)
	go kelilingPp(2, 3, kChan)

	var lChan = make(chan int)
	go luasPp(2, 3, lChan)

	var vChan = make(chan int)
	go volumeBalok(2, 3, 4, vChan)

	for i := 0; i < 3; i++ {
		select {
		case k := <-kChan:
			fmt.Println("Keliling PP:", k)
		case l := <-lChan:
			fmt.Println("Luas PP:", l)
		case v := <-vChan:
			fmt.Println("Volume Balok:", v)
		}
	}
}
