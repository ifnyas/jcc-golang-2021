package session

import (
	"context"
	"database/sql"
	"fmt"
	"jcc-golang-2021/util"
	"log"
	"time"
)

type Session struct {
	ID           int       `json:"id"`
	Courier      string    `json:"courier"`
	Note         string    `json:"note"`
	DeliveryCost float64   `json:"delivery_cost"`
	UserId       int       `json:"user_id"`
	StatusId     int       `json:"status_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

const (
	table = "t_session"
)

func PostDb(ctx context.Context, item Session) error {
	queryText := fmt.Sprintf(
		`INSERT INTO %v (
			courier, note, delivery_cost, user_id, status_id,
			created_at, updated_at
		) values('%v','%v', %v, %v, %v, NOW(), NOW())`,
		table,
		item.Courier,
		item.Note,
		item.DeliveryCost,
		item.UserId,
		item.StatusId,
	)
	_, err := util.ExecDb(ctx, queryText)
	return err
}

func PutDb(ctx context.Context, item Session) error {
	queryText := fmt.Sprintf(
		`UPDATE %v set 
		courier = '%v',
		note = '%v',
		delivery_cost = %v,
		status_id = %v,
		updated_at = NOW() 
		where id = %v`,
		table,
		item.Courier,
		item.Note,
		item.DeliveryCost,
		item.StatusId,
		item.ID,
	)
	_, err := util.ExecDb(ctx, queryText)
	return err
}

func GetByUserIdDb(ctx context.Context, productId int) ([]Session, error) {
	// get rows
	queryText := fmt.Sprintf(
		"SELECT * FROM %v WHERE user_id = %v",
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

func GetByIdDb(ctx context.Context, id int) ([]Session, error) {
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

func rowQueryNext(rowQuery *sql.Rows) []Session {
	var items []Session
	for rowQuery.Next() {
		var item Session
		var err error
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&item.ID,
			&item.Courier,
			&item.Note,
			&item.DeliveryCost,
			&item.UserId,
			&item.StatusId,
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
