package util

import (
	"context"
	"database/sql"
	"encoding/json"
	"jcc-golang-2021/config"
	"log"
	"net/http"
)

const (
	LayoutDateTime = "2006-01-02 15:04:05"
)

func ResponseJSON(w http.ResponseWriter, p interface{}, status int) {
	encoded, err := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, "Oops...", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(encoded))
}

func ExecDb(ctx context.Context, q string) (sql.Result, error) {
	// connect to sql
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	// send query
	res, err := db.ExecContext(ctx, q)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func QueryDb(ctx context.Context, q string) (*sql.Rows, error) {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}
	return db.QueryContext(ctx, q)
}
