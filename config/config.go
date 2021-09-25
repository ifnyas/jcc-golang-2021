package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// var (
// user string = os.Getenv("DB_USER")
// pass string = os.Getenv("DB_PASS")
// data string = os.Getenv("DB_DATA")
// host string = os.Getenv("DB_HOST")
// )

const (
	user string = "root"
	pass string = "Iqbal347"
	data string = "db_golang_yasir"
	host string = "tcp(117.53.47.152:3306)"
)

// HubToMySQL
func MySQL() (*sql.DB, error) {
	dsn := fmt.Sprintf("%v:%v@%v/%v", user, pass, host, data)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
