package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"jcc-golang-2021/util"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	RoleId    int       `json:"role_id"`
	FullName  string    `json:"full_name"`
	BirthDate string    `json:"birth_date"`
	ImageUrl  string    `json:"image_url"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	IsActive  int       `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const (
	table = "t_user"
)

func GetByBasicAuth(ctx context.Context, r *http.Request) []User {
	// check user and pass
	userAuth, passAuth, okAuth := r.BasicAuth()
	theUser, err := GetByUsernameDb(ctx, userAuth)
	if err == nil && theUser != nil && okAuth {
		passDecryptErr := bcrypt.CompareHashAndPassword(
			[]byte(theUser[0].Password),
			[]byte(passAuth))

		if passDecryptErr != nil {
			return nil
		}
	}
	return theUser
}

func PostDb(ctx context.Context, item User) error {
	queryText := fmt.Sprintf(
		"INSERT INTO %v ("+
			"username, `password`, role_id, full_name, birth_date,"+
			"image_url, phone, email, `address`, is_active, created_at, updated_at) "+
			"values('%v','%v', %v,'%v','%v','%v','%v','%v','%v', 1, NOW(), NOW())",
		table,
		item.Username,
		item.Password,
		item.RoleId,
		item.FullName,
		item.BirthDate,
		item.ImageUrl,
		item.Phone,
		item.Email,
		item.Address,
	)
	_, err := util.ExecDb(ctx, queryText)
	return err
}

func PutDb(ctx context.Context, item User) error {
	queryText := fmt.Sprintf(
		"UPDATE %v set "+
			"full_name = '%v',"+
			"birth_date = '%v',"+
			"image_url = '%v',"+
			"phone = '%v',"+
			"email = '%v',"+
			"`address` = '%v',"+
			"updated_at = NOW() "+
			"where id = %v",
		table,
		item.FullName,
		item.BirthDate,
		item.ImageUrl,
		item.Phone,
		item.Email,
		item.Address,
		item.ID,
	)
	_, err := util.ExecDb(ctx, queryText)
	return err
}

func ResetPassDb(ctx context.Context, item User) error {
	queryText := fmt.Sprintf(
		"UPDATE %v set `password` = '%v', updated_at = NOW() where id = %v",
		table,
		item.Password,
		item.ID,
	)
	_, err := util.ExecDb(ctx, queryText)
	return err
}

func DeactivateDb(ctx context.Context, id int) error {
	queryText := fmt.Sprintf(
		"UPDATE %v set is_active = 0, updated_at = NOW() where id = %v",
		table,
		id,
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

func GetByUsernameDb(ctx context.Context, username string) ([]User, error) {
	// get rows
	queryText := fmt.Sprintf(
		"SELECT * FROM %v WHERE username = '%v'",
		table,
		username)

	rowQuery, err := util.QueryDb(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}

	// parse rows
	var items []User
	for rowQuery.Next() {
		var item User
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&item.ID,
			&item.Username,
			&item.Password,
			&item.FullName,
			&item.BirthDate,
			&item.ImageUrl,
			&item.Phone,
			&item.Email,
			&item.Address,
			&item.IsActive,
			&item.RoleId,
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

func GetByIdDb(ctx context.Context, id int) ([]User, error) {
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
	var items []User
	for rowQuery.Next() {
		var item User
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&item.ID,
			&item.Username,
			&item.Password,
			&item.FullName,
			&item.BirthDate,
			&item.ImageUrl,
			&item.Phone,
			&item.Email,
			&item.Address,
			&item.IsActive,
			&item.RoleId,
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

func IsBasicAuthValid(rule int, userIdNeeded int, r *http.Request, ctx context.Context) bool {
	// init
	isAuthValid := false
	isUserValid := false

	// check user and pass
	userAuth, passAuth, okAuth := r.BasicAuth()
	theUser, err := GetByUsernameDb(ctx, userAuth)
	if err == nil && theUser != nil && okAuth {
		passDecrypt := bcrypt.CompareHashAndPassword(
			[]byte(theUser[0].Password),
			[]byte(passAuth))

		if passDecrypt == nil {
			isUserValid = true
		}

		switch rule {
		case 2: // creator
			if theUser[0].ID == userIdNeeded {
				isAuthValid = true
			}
			fallthrough
		default: // admin
			if theUser[0].RoleId == 1 {
				isAuthValid = true
			}
		}
	}
	return isAuthValid && isUserValid
}
