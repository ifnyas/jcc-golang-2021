package shop

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"jcc-golang-2021/util"
	"log"
	"time"
)

type Shop struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Tag       string    `json:"tag"`
	Name      string    `json:"name"`
	Detail    string    `json:"detail"`
	ImageUrl  string    `json:"image_url"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	IsActive  int       `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const (
	table = "t_shop"
)

func PostDb(ctx context.Context, item Shop) error {
	queryText := fmt.Sprintf(
		`INSERT INTO %v (
			'user_id', tag, 'name', detail,
			image_url, phone, email, 'address',
			is_active, created_at, updated_at
		) values(%v, '%v','%v', '%v', '%v', '%v', '%v', '%v', %v, NOW(), NOW())`,
		table,
		item.UserId,
		item.Tag,
		item.Name,
		item.Detail,
		item.ImageUrl,
		item.Phone,
		item.Email,
		item.Address,
		item.IsActive,
	)
	_, err := util.ExecDb(ctx, queryText)
	return err
}

func PutDb(ctx context.Context, item Shop) error {
	queryText := fmt.Sprintf(
		`UPDATE %v set 
		'user_id' = %v,
		'name' = '%v',
		detail = '%v',
		image_url = '%v',
		phone = '%v',
		email = '%v',
		'address' = '%v',
		is_active = %v,
		updated_at = NOW() 
		where id = %v`,
		table,
		item.UserId,
		item.Name,
		item.Detail,
		item.ImageUrl,
		item.Phone,
		item.Email,
		item.Address,
		item.IsActive,
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

func GetByTagDb(ctx context.Context, username string) ([]Shop, error) {
	// get rows
	queryText := fmt.Sprintf(
		"SELECT * FROM %v WHERE tag = '%v'",
		table,
		username)

	rowQuery, err := util.QueryDb(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}

	// parse rows
	var items []Shop
	for rowQuery.Next() {
		var item Shop
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&item.ID,
			&item.Tag,
			&item.Name,
			&item.Detail,
			&item.ImageUrl,
			&item.Phone,
			&item.Email,
			&item.Address,
			&item.IsActive,
			&item.UserId,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
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
	return items, nil
}

func GetByIdDb(ctx context.Context, id int) ([]Shop, error) {
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
	var items []Shop
	for rowQuery.Next() {
		var item Shop
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&item.ID,
			&item.Tag,
			&item.Name,
			&item.Detail,
			&item.ImageUrl,
			&item.Phone,
			&item.Email,
			&item.Address,
			&item.IsActive,
			&item.UserId,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
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
	return items, nil
}
