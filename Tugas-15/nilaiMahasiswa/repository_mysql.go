package nilaiMahasiswa

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
	"tugas-15/config"
	"tugas-15/models"
)

const (
	table          = "nilai_mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll Movie
func GetAll(ctx context.Context) ([]models.NilaiMahasiswa, error) {

	var scores []models.NilaiMahasiswa

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var score models.NilaiMahasiswa
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&score.ID,
			&score.Nama,
			&score.MataKuliah,
			&score.Nilai,
			&score.IndeksNilai,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		score.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		score.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		scores = append(scores, score)
	}

	return scores, nil
}

// Insert Movie
func Insert(ctx context.Context, score models.NilaiMahasiswa) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (nama, mata_kuliah, nilai, indeks_nilai, created_at, updated_at) values('%v', '%v',%v, '%v', NOW(), NOW())", table,
		score.Nama,
		score.MataKuliah,
		score.Nilai,
		score.IndeksNilai,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Movie
func Update(ctx context.Context, score models.NilaiMahasiswa, id string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set nama ='%s', mata_kuliah = '%s',  nilai = %d,  indeks_nilai = '%s', updated_at = NOW() where id = %s",
		table,
		score.Nama,
		score.MataKuliah,
		score.Nilai,
		score.IndeksNilai,
		id,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Movie
func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table, id)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
