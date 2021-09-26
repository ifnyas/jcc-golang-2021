package route

import (
	"context"
	"jcc-golang-2021/model/product"
	"jcc-golang-2021/model/review"
	"jcc-golang-2021/model/shop"
	"jcc-golang-2021/model/user"
	"jcc-golang-2021/util"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func PostReview(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// parse form
	var item review.Review
	item.Note = r.PostFormValue("note")
	item.MediaUrl = r.PostFormValue("media_url")

	formRating, err := strconv.Atoi(r.PostFormValue("rating"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	if formRating > 5 {
		formRating = 5
	}
	if formRating < 1 {
		formRating = 1
	}
	item.Rating = formRating

	formProductId, err := strconv.Atoi(r.PostFormValue("product_id"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	item.ProductId = formProductId

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
	if !user.IsBasicAuthValid(2, item.UserId, r, ctx) {
		err := map[string]string{
			"status": "Unauthorized!",
		}
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// post data
	if err := review.PostDb(ctx, item); err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Success!",
	}
	util.ResponseJSON(w, res, http.StatusOK)
}

func PutReview(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get item from db
	itemId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	theItem, err := review.GetByIdDb(ctx, itemId)
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
	formMediaUrl := r.PostFormValue("media_url")

	formRating := r.PostFormValue("rating")
	formRatingInt, err := strconv.Atoi(formRating)
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	// update value
	if formNote != "" {
		theItem[0].Note = formNote
	}
	if formMediaUrl != "" {
		theItem[0].MediaUrl = formMediaUrl
	}
	if formRating != "" {
		if formRatingInt > 5 {
			formRatingInt = 5
		}
		if formRatingInt < 1 {
			formRatingInt = 1
		}
		theItem[0].Rating = formRatingInt
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
	if err := review.PutDb(ctx, theItem[0]); err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Success!",
	}
	util.ResponseJSON(w, res, http.StatusOK)
}

func PutResponse(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get item from db
	itemId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	theItem, err := review.GetByIdDb(ctx, itemId)
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
	formResponse := r.PostFormValue("response")

	// update value
	if formResponse != "" {
		theItem[0].Response = formResponse
	}

	// get user id from the shop owner
	productId := theItem[0].ProductId
	theProduct, err := product.GetByIdDb(ctx, productId)
	if theProduct == nil {
		err := map[string]string{
			"status": "Product not found!",
		}
		util.ResponseJSON(w, err, http.StatusNotFound)
		return
	}
	if err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	theShopId := theProduct[0].ShopId
	theShop, err := shop.GetByIdDb(ctx, theShopId)
	if theShop == nil {
		err := map[string]string{
			"status": "Shop not found!",
		}
		util.ResponseJSON(w, err, http.StatusNotFound)
		return
	}
	if err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	// check auth
	if !user.IsBasicAuthValid(2, theShop[0].UserId, r, ctx) {
		err := map[string]string{
			"status": "Unauthorized!",
		}
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// put db
	if err := review.PutDb(ctx, theItem[0]); err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Success!",
	}
	util.ResponseJSON(w, res, http.StatusOK)
}

func GetReviewAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get param
	var productIdInt int
	productIdParam := r.URL.Query().Get("product_id")

	if productIdParam == "" {
		productIdInt = -1
	} else {
		var err error
		productIdInt, err = strconv.Atoi(productIdParam)
		if err != nil {
			util.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}
	}

	// query
	items, err := review.GetByIdDb(ctx, productIdInt)
	if err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	// result
	util.ResponseJSON(w, items, http.StatusOK)
}

func GetReview(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get id
	itemId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		itemId = -1
	}

	// query
	items, err := review.GetByIdDb(ctx, itemId)
	if items == nil {
		util.ResponseJSON(w, err, http.StatusNotFound)
		return
	}
	if err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	// result
	util.ResponseJSON(w, items, http.StatusOK)
}
