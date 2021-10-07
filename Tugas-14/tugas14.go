package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	ID          uint   `json:"id"`
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mataKuliah"`
	Nilai       uint   `json:"Nilai"`
	IndeksNilai string `json:"indeksNilai"`
}

var Scores = []NilaiMahasiswa{}

func getIndeksNilai(angka int) (nilai string) {
	if angka >= 80 {
		nilai = "A"
	} else if angka >= 70 && angka < 80 {
		nilai = "B"
	} else if angka >= 60 && angka < 70 {
		nilai = "C"
	} else if angka >= 50 && angka < 60 {
		nilai = "D"
	} else {
		nilai = "E"
	}
	return
}

// Fungi Log yang berguna sebagai middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

// GetScores
func GetScores(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataScores, err := json.Marshal(Scores)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataScores)
		return
	}

	http.Error(w, "NOT FOUND", http.StatusNotFound)
}

// PostScore
func PostScore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var validation = false
	var Scr NilaiMahasiswa
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Scr); err != nil {
				log.Fatal(err)
			}
			if Scr.Nilai > 100 || Scr.Nilai < 0 {
				validation = true
				http.Error(w, "nilai minimal 0 dan maksimal 100", http.StatusBadRequest)
			} else {
				Scr.IndeksNilai = getIndeksNilai(int(Scr.Nilai))
			}
		} else {
			// parse dari form
			getID := r.PostFormValue("id")
			id, _ := strconv.Atoi(getID)
			nama := r.PostFormValue("nama")
			mataKuliah := r.PostFormValue("mataKuliah")
			getNilai := r.PostFormValue("nilai")
			nilai, _ := strconv.Atoi(getNilai)
			if nilai > 100 || nilai < 0 {
				validation = true
				http.Error(w, "nilai minimal 0 dan maksimal 100", http.StatusBadRequest)
			} else {
				Scr = NilaiMahasiswa{
					ID:          uint(id),
					Nama:        nama,
					MataKuliah:  mataKuliah,
					Nilai:       uint(nilai),
					IndeksNilai: getIndeksNilai(nilai),
				}
			}
		}

		if !validation {
			Scores = append(Scores, Scr)
			dataScores, _ := json.Marshal(Scr) // to byte
			w.Write(dataScores)                // cetak di browser
		}
		return
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return
}

func main() {
	http.HandleFunc("/scores", GetScores)
	http.Handle("/post_score", Auth(http.HandlerFunc(PostScore)))
	fmt.Println("server running at http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
