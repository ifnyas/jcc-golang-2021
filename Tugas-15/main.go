package main

import (
	"context"
	"encoding/json"
	"fmt"
	"jcc-golang-2021/Tugas-15/config"
	"jcc-golang-2021/Tugas-15/model"
	"jcc-golang-2021/Tugas-15/nilai"
	"jcc-golang-2021/Tugas-15/util"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const (
	loginUser = "user"
	loginPass = "pass"
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

	// router
	router := httprouter.New()
	http.Handle("/", auth(http.HandlerFunc(nilaiRoute)))

	// serve
	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func nilaiRoute(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "POST":
		{
			PostNilai(w, r, nil)
		}
	case r.Method == "GET":
		{
			GetNilai(w, r, nil)
		}
	default:
		w.Write([]byte("Fungsi hanya mendukung metode GET dan POST"))
	}
}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// post auth
		if r.Method == "POST" {
			// basic auth
			user, pass, ok := r.BasicAuth()

			// auth not ok
			if !ok {
				w.Write([]byte("Username atau Password tidak boleh kosong"))
				return
			}

			// input invalid
			if user != loginUser || pass != loginPass {
				w.Write([]byte("Username atau Password tidak sesuai"))
				return
			}
		}

		// input ok
		next.ServeHTTP(w, r)
	})
}

func GetNilai(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// init context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get
	nilaiAll, err := nilai.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}
	util.ResponseJSON(w, nilaiAll, http.StatusOK)
}

func PostNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// init context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// init val
	var item model.NilaiMahasiswa

	// get input
	if r.Header.Get("Content-Type") == "application/json" {
		type nilaiInput struct {
			Nama, MataKuliah string
			Nilai            int
		}
		var inputStructed nilaiInput
		json.NewDecoder(r.Body).Decode(&inputStructed)

		item.Nama = inputStructed.Nama
		item.MataKuliah = inputStructed.MataKuliah
		item.Nilai = uint(inputStructed.Nilai)
	} else {
		item.Nama = r.PostFormValue("Nama")
		item.MataKuliah = r.PostFormValue("MataKuliah")
		nilaiValue, _ := strconv.Atoi(r.PostFormValue("Nilai"))
		item.Nilai = uint(nilaiValue)
	}

	// mod
	if item.Nilai > 100 {
		item.Nilai = 100
	}
	item.IndeksNilai = model.GetIndeks(item.Nilai)

	// post
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	if err := nilai.Insert(ctx, item); err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}

	util.ResponseJSON(w, res, http.StatusCreated)
}
