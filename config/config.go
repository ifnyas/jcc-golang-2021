package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
//user string = os.Getenv("DB_USER")
//pass string = os.Getenv("DB_PASS")
//data string = os.Getenv("DB_DATA")
//host string = os.Getenv("DB_HOST")
)

const (
	testUser string = "root"
	testPass string = "Iqbal347"
	testData string = "db_golang_yasir"
	testHost string = "tcp(117.53.47.152:3306)"
)

// HubToMySQL
func MySQL() (*sql.DB, error) {
	dsn := fmt.Sprintf("%v:%v@%v/%v", testUser, testPass, testHost, testData)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
