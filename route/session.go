package route

import (
	"context"
	"jcc-golang-2021/model/session"
	"jcc-golang-2021/model/user"
	"jcc-golang-2021/util"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func PutSession(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get item from db
	itemId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	theItem, err := session.GetByIdDb(ctx, itemId)
	if theItem == nil {
		err := map[string]string{
			"status": "Item not found!",
		}
		util.ResponseJSON(w, err, http.StatusNotFound)
		return
	}
	if err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	// parse form
	formCourier := r.PostFormValue("courier")
	formNote := r.PostFormValue("note")

	formDeliveryCost := r.PostFormValue("delivery_cost")
	formDeliveryCostInt, err := strconv.ParseFloat(formDeliveryCost, 64)
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	formStatusId := r.PostFormValue("status_id")
	formStatusIdInt, err := strconv.Atoi(formStatusId)
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	if formStatusIdInt == 0 {
		err := map[string]string{
			"status": "status_id tidak boleh 0!",
		}
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// update value
	if formCourier != "" {
		theItem[0].Courier = formCourier
	}
	if formNote != "" {
		theItem[0].Note = formNote
	}
	if formDeliveryCost != "" {
		theItem[0].DeliveryCost = formDeliveryCostInt
	}
	if formStatusId != "" {
		theItem[0].StatusId = formStatusIdInt
	}

	// check auth
	if !user.IsBasicAuthValid(2, theItem[0].UserId, r, ctx) {
		err := map[string]string{
			"status": "Unauthorized!",
		}
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// put db
	if err := session.PutDb(ctx, theItem[0]); err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Success!",
	}
	util.ResponseJSON(w, res, http.StatusOK)

	// create new session if statusId != 0
	if theItem[0].StatusId != 0 {
		newSession := session.Session{
			ID:           0,
			Courier:      "",
			Note:         "",
			DeliveryCost: 0,
			UserId:       theItem[0].UserId,
			StatusId:     0,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now()}
		session.PostDb(ctx, newSession)
	}
}

func GetSession(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get id
	itemId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		itemId = -1
	}

	// get param
	userIdParam := r.URL.Query().Get("user_id")
	if itemId != -1 {
		userIdParam = ""
	}

	// query
	var items []session.Session
	if userIdParam == "" {
		items, err = session.GetByIdDb(ctx, itemId)
		if items == nil {
			err := map[string]string{
				"status": "Review not found!",
			}
			util.ResponseJSON(w, err, http.StatusNotFound)
			return
		}
		if err != nil {
			util.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}
	} else {
		userIdInt, err := strconv.Atoi(userIdParam)
		if err != nil {
			util.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}
		items, err = session.GetByUserIdDb(ctx, userIdInt)
		if items == nil {
			err := map[string]string{
				"status": "Session not found!",
			}
			util.ResponseJSON(w, err, http.StatusNotFound)
			return
		}
		if err != nil {
			util.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	// check auth
	rule := 2
	if itemId < 0 && userIdParam == "" {
		rule = 1
	}
	if !user.IsBasicAuthValid(rule, items[0].UserId, r, ctx) {
		err := map[string]string{
			"status": "Unauthorized!",
		}
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// result
	util.ResponseJSON(w, items, http.StatusOK)
}

/*
func PostSession(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// parse form
	var item session.Session
	item.Courier = r.PostFormValue("courier")
	item.Note = r.PostFormValue("note")
	item.StatusId = 1

	formDeliveryCost, err := strconv.ParseFloat(r.PostFormValue("delivery_cost"), 64)
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	item.DeliveryCost = formDeliveryCost

	formUserId := r.PostFormValue("user_id")
	if formUserId != "" {
		formUserIdInt, err := strconv.Atoi(formUserId)
		if err != nil {
			util.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}
		item.UserId = formUserIdInt
	} else {
		theUser := user.GetByBasicAuth(ctx, r)
		if theUser != nil {
			item.UserId = theUser[0].ID
		} else {
			err := map[string]string{
				"status": "Unauthorized!",
			}
			util.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}
	}

	// check auth
	if user.IsBasicAuthValid(2, item.UserId, r, ctx) {
		err := map[string]string{
			"status": "Unauthorized!",
		}
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// post data
	if err := session.PostDb(ctx, item); err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Success!",
	}
	util.ResponseJSON(w, res, http.StatusOK)
}
*/
