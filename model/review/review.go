package review

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"jcc-golang-2021/util"
	"log"
	"time"
)

type Review struct {
	ID        int       `json:"id"`
	Note      string    `json:"note"`
	Response  string    `json:"response"`
	MediaUrl  string    `json:"media_url"`
	Rating    int       `json:"rating"`
	UserId    int       `json:"user_id"`
	ProductId int       `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const (
	table = "t_review"
)

func PostDb(ctx context.Context, item Review) error {
	queryText := fmt.Sprintf(
		"INSERT INTO %v ("+
			"note, media_url, rating, user_id,"+
			"created_at, updated_at"+
			") values('%v','%v', %v, %v, NOW(), NOW())",
		table,
		item.Note,
		item.MediaUrl,
		item.Rating,
		item.UserId,
	)
	_, err := util.ExecDb(ctx, queryText)
	return err
}

func PutDb(ctx context.Context, item Review) error {
	queryText := fmt.Sprintf(
		"UPDATE %v set "+
			"note = '%v',"+
			"media_url = '%v',"+
			"rating = %v,"+
			"updated_at = NOW()"+
			"where id = %v",
		table,
		item.Note,
		item.MediaUrl,
		item.Rating,
		item.ID,
	)
	_, err := util.ExecDb(ctx, queryText)
	return err
}

func PutResponseDb(ctx context.Context, item Review) error {
	queryText := fmt.Sprintf(
		"UPDATE %v set response = '%v', updated_at = NOW() where id = %v",
		table,
		item.Response,
		item.ID,
	)
	_, err := util.ExecDb(ctx, queryText)
	return err
}

func DelDb(ctx context.Context, id int) error {
	queryText := fmt.Sprintf("DELETE FROM %v where id = %v", table, id)
	res, err := util.ExecDb(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := res.RowsAffected()
	if check == 0 {
		return errors.New("item not found")
	}
	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

func GetByProductIdDb(ctx context.Context, productId int) ([]Review, error) {
	// get rows
	queryText := fmt.Sprintf(
		"SELECT * FROM %v WHERE product_id = %v",
		table,
		productId)

	rowQuery, err := util.QueryDb(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}

	// parse rows
	items := rowQueryNext(rowQuery)
	return items, nil
}

func GetByIdDb(ctx context.Context, id int) ([]Review, error) {
	// get rows
	queryText := ""
	if id < 0 {
		queryText = fmt.Sprintf(
			"SELECT * FROM %v",
			table)
	} else {
		queryText = fmt.Sprintf(
			"SELECT * FROM %v WHERE id = %v",
			table, id)
	}

	rowQuery, err := util.QueryDb(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}

	// parse rows
	items := rowQueryNext(rowQuery)
	return items, nil
}

func rowQueryNext(rowQuery *sql.Rows) []Review {
	var items []Review
	for rowQuery.Next() {
		var item Review
		var err error
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&item.ID,
			&item.Note,
			&item.Response,
			&item.MediaUrl,
			&item.Rating,
			&item.UserId,
			&item.ProductId,
			&createdAt,
			&updatedAt); err != nil {
			log.Fatal(err)
		}

		item.CreatedAt, err = time.Parse(util.LayoutDateTime, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		item.UpdatedAt, err = time.Parse(util.LayoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		items = append(items, item)
	}
	return items
}
