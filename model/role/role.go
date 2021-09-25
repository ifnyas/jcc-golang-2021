package role

import (
	"context"
	"database/sql"
	"fmt"
	"jcc-golang-2021/util"
	"log"
	"time"
)

const (
	table = "t_role"
)

// title: admin | customer | seller
type Role struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetByIdDb(ctx context.Context, id int) ([]Role, error) {
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

func rowQueryNext(rowQuery *sql.Rows) []Role {
	var items []Role
	for rowQuery.Next() {
		var item Role
		var err error
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&item.ID,
			&item.Title,
			&item.Detail,
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
