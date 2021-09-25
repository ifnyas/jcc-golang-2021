package route

import (
	"context"
	"jcc-golang-2021/model/cart"
	"jcc-golang-2021/model/session"
	"jcc-golang-2021/model/user"
	"jcc-golang-2021/util"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func PostCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// parse form
	var item cart.Cart
	item.ProductName = r.PostFormValue("product_name")
	item.Note = r.PostFormValue("note")

	formProductPrice, err := strconv.ParseFloat(r.PostFormValue("product_price"), 64)
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	item.ProductPrice = formProductPrice

	formProductPriceMod, err := strconv.ParseFloat(r.PostFormValue("product_price_mod"), 64)
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	item.ProductPriceMod = formProductPriceMod

	formAmount, err := strconv.Atoi(r.PostFormValue("amount"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	item.Amount = formAmount

	formSessionId, err := strconv.Atoi(r.PostFormValue("session_id"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	item.SessionId = formSessionId

	// check auth
	if !isCartBasicAuthValid(r, ctx, 2, item.SessionId) {
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// post data
	if err := cart.PostDb(ctx, item); err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Success!",
	}
	util.ResponseJSON(w, res, http.StatusOK)
}

func PutCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get item from db
	itemId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	theItem, err := cart.GetByIdDb(ctx, itemId)
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
	formNote := r.PostFormValue("note")

	formProductPriceMod := r.PostFormValue("product_price_mod")
	formProductPriceModInt, err := strconv.ParseFloat(formProductPriceMod, 64)
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	formAmount := r.PostFormValue("amount")
	formAmountInt, err := strconv.Atoi(formAmount)
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	// update value
	if formNote != "" {
		theItem[0].Note = formNote
	}
	if formProductPriceMod != "" {
		theItem[0].ProductPriceMod = formProductPriceModInt
	}
	if formAmount != "" {
		theItem[0].Amount = formAmountInt
	}

	// check auth
	if !isCartBasicAuthValid(r, ctx, 2, theItem[0].SessionId) {
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// put db
	if err := cart.PutDb(ctx, theItem[0]); err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Success!",
	}
	util.ResponseJSON(w, res, http.StatusOK)
}

func DelCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get item from db
	itemId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	theItem, err := cart.GetByIdDb(ctx, itemId)
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

	// check auth
	if !isCartBasicAuthValid(r, ctx, 2, theItem[0].SessionId) {
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// exec
	if err := cart.DelDb(ctx, itemId); err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Success!",
	}
	util.ResponseJSON(w, res, http.StatusOK)
}

func GetCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get id
	itemId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		itemId = -1
	}

	// get param
	sessionIdParam := r.URL.Query().Get("session_id")
	if itemId != -1 {
		sessionIdParam = ""
	}

	// query
	var items []cart.Cart
	if sessionIdParam == "" {
		items, err = cart.GetByIdDb(ctx, itemId)
		if items == nil {
			err := map[string]string{
				"status": "Cart not found!",
			}
			util.ResponseJSON(w, err, http.StatusNotFound)
			return
		}
		if err != nil {
			util.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}
	} else {
		sessionIdInt, err := strconv.Atoi(sessionIdParam)
		if err != nil {
			util.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}
		items, err = cart.GetBySessionIdDb(ctx, sessionIdInt)
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
	if itemId < 0 && sessionIdParam == "" {
		rule = 1
	}
	if !isCartBasicAuthValid(r, ctx, rule, items[0].SessionId) {
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// result
	util.ResponseJSON(w, items, http.StatusOK)
}

func isCartBasicAuthValid(r *http.Request, ctx context.Context, rule int, sessionId int) bool {
	isValid := false
	theSession, err := session.GetByIdDb(ctx, sessionId)
	isBasicAuthValid := user.IsBasicAuthValid(rule, theSession[0].UserId, r, ctx)
	if theSession != nil && err == nil && isBasicAuthValid {
		isValid = true
	}
	return isValid
}
