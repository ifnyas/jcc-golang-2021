package nilai

import (
	"context"
	"fmt"
	"jcc-golang-2021/Tugas-15/config"
	"jcc-golang-2021/Tugas-15/model"
	"log"
	"time"
)

const (
	table          = "nilai_mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

// Store Nilai
func Insert(ctx context.Context, nilai model.NilaiMahasiswa) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf(
		"INSERT INTO %v (nama, mata_kuliah, indeks_nilai, nilai, created_at, updated_at) values('%v','%v','%v',%v, NOW(), NOW())",
		table,
		nilai.Nama,
		nilai.MataKuliah,
		nilai.IndeksNilai,
		nilai.Nilai)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// GetAll
func GetAll(ctx context.Context) ([]model.NilaiMahasiswa, error) {
	var nilaiAll []model.NilaiMahasiswa
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By created_at DESC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var nilai model.NilaiMahasiswa
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&nilai.ID,
			&nilai.Nama,
			&nilai.MataKuliah,
			&nilai.IndeksNilai,
			&nilai.Nilai,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		nilai.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		nilai.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		nilaiAll = append(nilaiAll, nilai)
	}
	return nilaiAll, nil
}
