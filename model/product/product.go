package product

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"jcc-golang-2021/util"
	"log"
	"time"
)

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Detail    string    `json:"detail"`
	Category  string    `json:"category"`
	ImageUrl  string    `json:"image_url"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	ShopId    int       `json:"shop_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const (
	table = "t_product"
)

func PostDb(ctx context.Context, item Product) error {
	queryText := fmt.Sprintf(
		"INSERT INTO %v ("+
			"shop_id, `name`, detail, price, category,"+
			"image_url, stock, created_at, updated_at"+
			") values(%v, '%v','%v', %v, '%v', '%v', %v, NOW(), NOW())",
		table,
		item.ShopId,
		item.Name,
		item.Detail,
		item.Price,
		item.Category,
		item.ImageUrl,
		item.Stock,
	)
	_, err := util.ExecDb(ctx, queryText)
	return err
}

func PutDb(ctx context.Context, item Product) error {
	queryText := fmt.Sprintf(
		"UPDATE %v set "+
			"shop_id = %v,"+
			"`name` = '%v',"+
			"detail = '%v',"+
			"price = %v,"+
			"category = '%v',"+
			"image_url = '%v',"+
			"stock = %v,"+
			"updated_at = NOW() "+
			"where id = %v",
		table,
		item.ShopId,
		item.Name,
		item.Detail,
		item.Price,
		item.Category,
		item.ImageUrl,
		item.Stock,
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

func GetByShopIdDb(ctx context.Context, shopId int) ([]Product, error) {
	// get rows
	queryText := fmt.Sprintf(
		"SELECT * FROM %v WHERE shop_id = %v",
		table,
		shopId)

	rowQuery, err := util.QueryDb(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}

	// parse rows
	var items []Product
	for rowQuery.Next() {
		var item Product
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&item.ID,
			&item.Name,
			&item.Detail,
			&item.Category,
			&item.ImageUrl,
			&item.Price,
			&item.Stock,
			&item.ShopId,
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

func GetByIdDb(ctx context.Context, id int) ([]Product, error) {
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
	var items []Product
	for rowQuery.Next() {
		var item Product
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&item.ID,
			&item.Name,
			&item.Detail,
			&item.Category,
			&item.ImageUrl,
			&item.Price,
			&item.Stock,
			&item.ShopId,
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
