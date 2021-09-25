package cart

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"jcc-golang-2021/util"
	"log"
	"time"
)

type Cart struct {
	ID              int       `json:"id"`
	ProductName     string    `json:"product_name"`
	Note            string    `json:"note"`
	ProductPrice    float64   `json:"product_price"`
	ProductPriceMod float64   `json:"product_price_mod"`
	Amount          int       `json:"amount"`
	SessionId       int       `json:"session_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

const (
	table = "t_cart"
)

func PostDb(ctx context.Context, item Cart) error {
	queryText := fmt.Sprintf(
		`INSERT INTO %v (
			product_name, note, product_price, product_price_mod,
			amount, session_id, created_at, updated_at
		) values('%v','%v', %v, %v, %v, %v, NOW(), NOW())`,
		table,
		item.ProductName,
		item.Note,
		item.ProductPrice,
		item.ProductPriceMod,
		item.Amount,
		item.SessionId,
	)
	_, err := util.ExecDb(ctx, queryText)
	return err
}

func PutDb(ctx context.Context, item Cart) error {
	queryText := fmt.Sprintf(
		`UPDATE %v set 
		note = '%v',
		product_price_mod = %v,
		amount = %v,
		updated_at = NOW() 
		where id = %v`,
		table,
		item.Note,
		item.ProductPriceMod,
		item.Amount,
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

func GetBySessionIdDb(ctx context.Context, sessionId int) ([]Cart, error) {
	// get rows
	queryText := fmt.Sprintf(
		"SELECT * FROM %v WHERE session_id = %v",
		table,
		sessionId)

	rowQuery, err := util.QueryDb(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}

	// parse rows
	items := rowQueryNext(rowQuery)
	return items, nil
}

func GetByIdDb(ctx context.Context, id int) ([]Cart, error) {
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

func rowQueryNext(rowQuery *sql.Rows) []Cart {
	var items []Cart
	for rowQuery.Next() {
		var item Cart
		var err error
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&item.ID,
			&item.ProductName,
			&item.Note,
			&item.ProductPrice,
			&item.ProductPriceMod,
			&item.Amount,
			&item.SessionId,
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
