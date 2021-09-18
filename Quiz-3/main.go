package main

import (
	"encoding/json"
	"fmt"
	"jcc-golang-2021/Quiz-3/config"
	"jcc-golang-2021/Quiz-3/model"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	authUser  = []string{"admin", "editor", "trainer"}
	authPass  = []string{"password", "secret", "rahasia"}
	booksTemp = []model.Book{}
)

func main() {
	// connect sql
	db, e := config.MySQL()
	if e != nil {
		log.Fatal(e)
	}
	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}
	fmt.Println("Success")

	// route
	bangunDatarRoute()
	booksRoute()

	// serve
	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func bangunDatarRoute() {
	http.Handle("/bangun-datar/segitiga-sama-sisi", auth(http.HandlerFunc(segitigaSamaSisiRoute)))
	http.Handle("/bangun-datar/persegi", auth(http.HandlerFunc(persegiRoute)))
	http.Handle("/bangun-datar/persegi-panjang", auth(http.HandlerFunc(persegiPanjangRoute)))
	http.Handle("/bangun-datar/lingkaran", auth(http.HandlerFunc(lingkaranRoute)))
	http.Handle("/bangun-datar/jajar-genjang", auth(http.HandlerFunc(jajarGenjangRoute)))

	// example for goroutine in persegi
	// example for channel in lingkaran
}

// books route
func booksRoute() {
	route := func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == "POST":
			PostBook(w, r)
		case r.Method == "GET":
			GetBook(w, r)
		case r.Method == "PUT":
			PutBook(w, r)
		case r.Method == "DELETE":
			DelBook(w, r)
		default:
			w.Write([]byte("Fungsi hanya mendukung metode POST/GET/PUT/DEL"))
		}
	}
	http.Handle("/books", auth(http.HandlerFunc(route)))
	http.Handle("/books/", auth(http.HandlerFunc(route)))
}

func PostBook(w http.ResponseWriter, r *http.Request) {
	var newBook model.Book

	// get input
	if r.Header.Get("Content-Type") == "application/json" {
		type BookInput struct {
			Title        string `json:"title"`
			Description  string `json:"description"`
			Image_url    string `json:"image_url"`
			Release_year int    `json:"release_year"`
			Price        int    `json:"price"`
			Total_page   string `json:"total_page"`
		}
		var bookStructed BookInput
		json.NewDecoder(r.Body).Decode(&bookStructed)

		newBook.Title = bookStructed.Title
		newBook.Description = bookStructed.Description
		newBook.Image_url = bookStructed.Image_url
		newBook.Total_page = bookStructed.Total_page
		newBook.Release_year = bookStructed.Release_year
		newBook.Price = model.GetPriceWithCurrency(bookStructed.Price)
	} else {
		newBook.Title = r.PostFormValue("title")
		newBook.Description = r.PostFormValue("description")
		newBook.Image_url = r.PostFormValue("image_url")
		newBook.Total_page = r.PostFormValue("total_page")
		newBook.Release_year, _ = strconv.Atoi(r.PostFormValue("release_year"))
		priceInput, _ := strconv.Atoi(r.PostFormValue("price"))
		newBook.Price = model.GetPriceWithCurrency(priceInput)
	}

	// validation
	errMsg := ""
	if !model.IsImageUrlValid(newBook.Image_url) {
		errMsg += "image_url tidak dapat diakses"
	}
	if !model.IsReleaseYearValid(newBook.Release_year) {
		if errMsg != "" {
			errMsg += " dan "
		}
		errMsg += "release_year harus di antara 1980 - 2021"
	}

	// mod
	latestIndex := 0
	for _, book := range booksTemp {
		latestIndex = book.ID
	}

	newBook.ID = latestIndex + 1
	tebal, _ := strconv.Atoi(newBook.Total_page)
	newBook.Kategori_ketebalan = model.GetKategoriKetebalan(tebal)
	newBook.CreatedAt = time.Now()
	newBook.UpdatedAt = time.Now()

	// push new value
	booksTemp = append(booksTemp, newBook)

	// show result
	if errMsg != "" {
		errJson := "{ \"error\" : \"" + errMsg + "\" }"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(errJson))
	} else {
		bookJson, _ := json.Marshal(newBook)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bookJson)
	}
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	result := []model.Book{}
	indexToDel := []int{}

	// find indexes to delete
	qTitle := r.URL.Query().Get("title")
	if qTitle != "" {
		for _, book := range booksTemp {
			if !strings.Contains(strings.ToLower(book.Title), strings.ToLower(qTitle)) {
				indexToDel = append(indexToDel, book.ID)
			}
		}
	}

	qMinYear := r.URL.Query().Get("minYear")
	if qMinYear != "" {
		qMinYearInt, _ := strconv.Atoi(qMinYear)
		for _, book := range booksTemp {
			if book.Release_year < qMinYearInt {
				indexToDel = append(indexToDel, book.ID)
			}
		}
	}

	qMaxYear := r.URL.Query().Get("maxYear")
	if qMaxYear != "" {
		qMaxYearInt, _ := strconv.Atoi(qMaxYear)
		for _, book := range booksTemp {
			if book.Release_year > qMaxYearInt {
				indexToDel = append(indexToDel, book.ID)
			}
		}
	}

	qMinPage := r.URL.Query().Get("minPage")
	if qMinPage != "" {
		qMinPageInt, _ := strconv.Atoi(qMinPage)
		for _, book := range booksTemp {
			totalPageInt, _ := strconv.Atoi(book.Total_page)
			if totalPageInt < qMinPageInt-1 {
				indexToDel = append(indexToDel, book.ID)
			}
		}
	}

	qMaxPage := r.URL.Query().Get("maxPage")
	if qMaxPage != "" {
		qMaxPageInt, _ := strconv.Atoi(qMaxPage)
		for _, book := range booksTemp {
			totalPageInt, _ := strconv.Atoi(book.Total_page)
			if totalPageInt > qMaxPageInt+1 {
				indexToDel = append(indexToDel, book.ID)
			}
		}
	}

	qSort := r.URL.Query().Get("sort")
	isAsc := false
	isDesc := false
	if qSort == "asc" {
		isAsc = true
	}
	if qSort == "desc" {
		isDesc = true
	}

	// remove duplicate indexes
	sort.Ints(indexToDel)
	uniqueIndexes := []int{}
	for i, n := range indexToDel {
		if i != 0 {
			if indexToDel[i] != indexToDel[i-1] {
				uniqueIndexes = append(uniqueIndexes, n)
			}
		} else {
			uniqueIndexes = append(uniqueIndexes, n)
		}

	}
	indexToDel = uniqueIndexes

	// append chosen
	for _, book := range booksTemp {
		toDel := false
		for i, n := range indexToDel {
			if book.ID == n {
				toDel = true
				indexToDel = append(indexToDel[:i], indexToDel[i+1:]...)
			}
			break
		}
		if !toDel {
			result = append(result, book)
		}
	}
	if isAsc {
		sort.Slice(result, func(i, j int) bool {
			return result[i].ID < result[j].ID
		})
	}
	if isDesc {
		sort.Slice(result, func(i, j int) bool {
			return result[i].ID > result[j].ID
		})
	}

	// show result
	bukuJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bukuJson)
}

func PutBook(w http.ResponseWriter, r *http.Request) {
	var newBook model.Book
	errMsg := ""
	newBook.ID, _ = strconv.Atoi(strings.Split(r.URL.Path, "/")[2])

	// get input
	if r.Header.Get("Content-Type") == "application/json" {
		type BookInput struct {
			Title        string `json:"title"`
			Description  string `json:"description"`
			Image_url    string `json:"image_url"`
			Release_year int    `json:"release_year"`
			Price        int    `json:"price"`
			Total_page   string `json:"total_page"`
		}
		var bookStructed BookInput
		json.NewDecoder(r.Body).Decode(&bookStructed)

		newBook.Title = bookStructed.Title
		newBook.Description = bookStructed.Description
		newBook.Image_url = bookStructed.Image_url
		newBook.Total_page = bookStructed.Total_page
		newBook.Release_year = bookStructed.Release_year
		newBook.Price = model.GetPriceWithCurrency(bookStructed.Price)
	} else {
		newBook.Title = r.PostFormValue("title")
		newBook.Description = r.PostFormValue("description")
		newBook.Image_url = r.PostFormValue("image_url")
		newBook.Total_page = r.PostFormValue("total_page")
		newBook.Release_year, _ = strconv.Atoi(r.PostFormValue("release_year"))
		priceInput, _ := strconv.Atoi(r.PostFormValue("price"))
		newBook.Price = model.GetPriceWithCurrency(priceInput)
	}

	chosenOneIndex := 0
	for i, book := range booksTemp {
		if book.ID == newBook.ID {
			chosenOneIndex = i
		}
	}

	if chosenOneIndex == 0 {
		errMsg = "Update failed, book not found!"
	} else {
		// validation
		if !model.IsImageUrlValid(newBook.Image_url) {
			errMsg += "image_url tidak dapat diakses"
		}
		if !model.IsReleaseYearValid(newBook.Release_year) {
			if errMsg != "" {
				errMsg += " dan "
			}
			errMsg += "release_year harus di antara 1980 - 2021"
		}

		// mod
		tebal, _ := strconv.Atoi(newBook.Total_page)
		newBook.Kategori_ketebalan = model.GetKategoriKetebalan(tebal)
		newBook.CreatedAt = booksTemp[newBook.ID-1].CreatedAt
		newBook.UpdatedAt = time.Now()

		// push new value
		for i, book := range booksTemp {
			if book.ID == newBook.ID {
				booksTemp[i] = newBook
			}
		}
	}

	// show result
	if errMsg != "" {
		errJson := "{ \"error\" : \"" + errMsg + "\" }"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(errJson))
	} else {
		bookJson, _ := json.Marshal(newBook)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bookJson)
	}
}

func DelBook(w http.ResponseWriter, r *http.Request) {
	resultMsg := "{ \"result\" : \"Delete failed, book not found!\" }"
	delId, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])

	for i, book := range booksTemp {
		if book.ID == delId {
			booksTemp = append(booksTemp[:i], booksTemp[i+1:]...)
			resultMsg = "{ \"result\" : \"Delete success!\" }"
		}
	}

	// show result
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resultMsg))
}

// bangun datar route
func segitigaSamaSisiRoute(w http.ResponseWriter, r *http.Request) {
	var datar model.BangunDatar

	hitung := r.URL.Query().Get("hitung")
	alas, err := strconv.Atoi(r.URL.Query().Get("alas"))
	if err != nil {
		alas = 0
	}
	tinggi, err := strconv.Atoi(r.URL.Query().Get("tinggi"))
	if err != nil {
		tinggi = 0
	}

	switch hitung {
	case "luas":
		datar.Result = float64(alas * tinggi / 2)
	case "keliling":
		datar.Result = float64(alas * 3)
	default:
		datar.Result = 0
	}

	printDatar(datar, w)
}

func persegiRoute(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	var datar model.BangunDatar

	printPersegi := func() {
		printDatar(datar, w)
		wg.Done()
	}

	hitung := r.URL.Query().Get("hitung")
	sisi, err := strconv.Atoi(r.URL.Query().Get("sisi"))
	if err != nil {
		sisi = 0
	}

	switch hitung {
	case "luas":
		datar.Result = math.Pow(float64(sisi), 2)
	case "keliling":
		datar.Result = float64(sisi * 4)
	default:
		datar.Result = 0
	}

	wg.Add(1)
	go printPersegi()
	wg.Wait()
}

func persegiPanjangRoute(w http.ResponseWriter, r *http.Request) {
	var datar model.BangunDatar

	hitung := r.URL.Query().Get("hitung")
	panjang, err := strconv.Atoi(r.URL.Query().Get("panjang"))
	if err != nil {
		panjang = 0
	}
	lebar, err := strconv.Atoi(r.URL.Query().Get("lebar"))
	if err != nil {
		lebar = 0
	}

	switch hitung {
	case "luas":
		datar.Result = float64(panjang * lebar)
	case "keliling":
		datar.Result = float64(2 * (panjang + lebar))
	default:
		datar.Result = 0
	}

	printDatar(datar, w)
}

func lingkaranRoute(w http.ResponseWriter, r *http.Request) {
	var datar model.BangunDatar

	hitung := r.URL.Query().Get("hitung")
	jariJari, err := strconv.Atoi(r.URL.Query().Get("jariJari"))
	if err != nil {
		jariJari = 0
	}

	inputChan := func(ch chan int, rs int) {
		ch <- rs
		close(ch)
	}

	jariJariChannel := make(chan int)
	go inputChan(jariJariChannel, jariJari)

	switch hitung {
	case "luas":
		datar.Result = math.Round(math.Pi * math.Pow(float64(<-jariJariChannel), 2))
	case "keliling":
		datar.Result = math.Round(2 * math.Pi * float64(<-jariJariChannel))
	default:
		datar.Result = 0
	}

	printDatar(datar, w)
}

func jajarGenjangRoute(w http.ResponseWriter, r *http.Request) {
	var datar model.BangunDatar

	hitung := r.URL.Query().Get("hitung")
	alas, err := strconv.Atoi(r.URL.Query().Get("alas"))
	if err != nil {
		alas = 0
	}
	tinggi, err := strconv.Atoi(r.URL.Query().Get("tinggi"))
	if err != nil {
		tinggi = 0
	}
	sisi, err := strconv.Atoi(r.URL.Query().Get("sisi"))
	if err != nil {
		sisi = 0
	}

	switch hitung {
	case "luas":
		datar.Result = float64(alas * tinggi)
	case "keliling":
		datar.Result = float64((2 * alas) + (2 * sisi))
	default:
		datar.Result = 0
	}

	printDatar(datar, w)
}

func printDatar(d model.BangunDatar, w http.ResponseWriter) {
	datarJson, _ := json.Marshal(d)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(datarJson)
}

// auth
func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// basic auth
		user, pass, ok := r.BasicAuth()

		// check which account used
		passSelected := ""
		isUserRegistered := false

		for i, u := range authUser {
			if u == user {
				passSelected = authPass[i]
				isUserRegistered = true
			}
		}

		// post auth
		if r.Method != "POST" {
			// auth not ok
			if !ok {
				w.Write([]byte("Username atau Password tidak boleh kosong"))
				return
			}

			// input invalid
			if !isUserRegistered || pass != passSelected {
				w.Write([]byte("Username atau Password tidak sesuai"))
				return
			}
		}

		// input ok
		next.ServeHTTP(w, r)
	})
}
