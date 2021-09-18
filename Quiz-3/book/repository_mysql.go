package nilai

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"jcc-golang-2021/Quiz-3/config"
	"jcc-golang-2021/Quiz-3/model"
	"log"
	"time"
)

const (
	table          = "t_book"
	layoutDateTime = "2006-01-02 15:04:05"
)

// Create
func Insert(ctx context.Context, book model.Book) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf(
		"INSERT INTO %v (title, description, image_url, release_year, price, total_page, kategori_ketebalan) values('%v','%v','%v',%v, '%v','%v','%v', NOW(), NOW())",
		table,
		book.Title,
		book.Description,
		book.Image_url,
		book.Release_year,
		book.Price,
		book.Total_page,
		book.Kategori_ketebalan)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Read
func GetAll(ctx context.Context) ([]model.Book, error) {
	var bookAll []model.Book
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var book model.Book
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&book.ID,
			&book.Title,
			&book.Description,
			&book.Image_url,
			&book.Release_year,
			&book.Price,
			&book.Total_page,
			&book.Kategori_ketebalan,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		book.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		book.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		bookAll = append(bookAll, book)
	}
	return bookAll, nil
}

// Update
func Update(ctx context.Context, book model.Book, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set title = '%s', description = '%s', image_url = '%s', release_year = %d, price = '%s', total_page = '%s', kategori_ketebalan = '%s', updated_at = NOW() where id = %d",
		table,
		book.Title,
		book.Description,
		book.Image_url,
		book.Release_year,
		book.Price,
		book.Total_page,
		book.Kategori_ketebalan,
		book.ID)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

// DELETE
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
		return errors.New("id not found")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
