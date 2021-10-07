package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tugas-15/models"
	"tugas-15/nilaiMahasiswa"
	"tugas-15/utils"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	router.GET("/nilai_mahasiswa", GetScores)
	router.POST("/nilai_mahasiswa/create", PostScore)
	router.PUT("/nilai_mahasiswa/:id/update", UpdateScore)
	router.DELETE("/nilai_mahasiswa/:id/delete", DeleteScore)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

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

// Read
// GetScores
func GetScores(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	allScores, err := nilaiMahasiswa.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, allScores, http.StatusOK)
}

// Create
// PostScore
func PostScore(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var score models.NilaiMahasiswa

	if err := json.NewDecoder(r.Body).Decode(&score); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	score.IndeksNilai = getIndeksNilai(int(score.Nilai))

	if err := nilaiMahasiswa.Insert(ctx, score); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateScore
func UpdateScore(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var score models.NilaiMahasiswa

	if err := json.NewDecoder(r.Body).Decode(&score); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idScore = ps.ByName("id")

	score.IndeksNilai = getIndeksNilai(int(score.Nilai))

	if err := nilaiMahasiswa.Update(ctx, score, idScore); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteScore
func DeleteScore(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idScore = ps.ByName("id")

	if err := nilaiMahasiswa.Delete(ctx, idScore); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}
