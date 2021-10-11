package route

import (
	"fmt"
	"jcc-golang-2021/db"
	"jcc-golang-2021/util"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const (
	radius = 5000
)

func Search(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// init params
	var params = []string{"lat", "lng", "category_id", "page", "limit"}
	var lat, lng, minLat, maxLat, minLng, maxLng float64
	queries := r.URL.RawQuery

	// parse params
	for _, param := range params {
		// init val
		var err error
		var key = r.URL.Query().Get(param)

		// parse
		if key != "" {
			switch param {
			case "lat":
				lat, err = strconv.ParseFloat(key, 64)
			case "lng":
				lng, err = strconv.ParseFloat(key, 64)
			case "category_id":
				_, err = strconv.Atoi(key)
			case "page":
				_, err = strconv.Atoi(key)
			case "limit":
				_, err = strconv.Atoi(key)
			}
		}

		// err handler
		if err != nil {
			util.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}
	}

	// get loc between
	locBetween := util.LocBetween(lat, lng, radius)
	if lat != 0.0 {
		minLat = locBetween[0]
		maxLat = locBetween[1]
		queries += "&min_lat=" +
			fmt.Sprintf("%f", minLat) +
			"&max_lat=" +
			fmt.Sprintf("%f", maxLat)
	}
	if lng != 0.0 {
		minLng = locBetween[2]
		maxLng = locBetween[3]
		queries += "&min_lng=" +
			fmt.Sprintf("%f", minLng) +
			"&max_lng=" +
			fmt.Sprintf("%f", maxLng)
	}

	// query
	items := db.GetData(w, queries)

	// result
	util.ResponseJSON(w, items, http.StatusOK)
}
