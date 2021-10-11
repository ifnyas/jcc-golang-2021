package db

import (
	"encoding/json"
	"io/ioutil"
	"jcc-golang-2021/util"
	"net/http"
	"os"
)

var (
	dbUrl = os.Getenv("DB_SHEET")
)

func GetData(w http.ResponseWriter, queries string) interface{} {
	// init data
	var data interface{}

	// request
	res, err := http.Get(dbUrl + "?" + queries)
	util.ErrHandler(err)

	// read body
	body, err := ioutil.ReadAll(res.Body)
	util.ErrHandler(err)

	// parse json
	err = json.Unmarshal(body, &data)
	util.ErrHandler(err)

	// return
	return data
}
