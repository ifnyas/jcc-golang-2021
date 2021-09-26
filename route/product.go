package route

import (
	"context"
	"fmt"
	"jcc-golang-2021/model/product"
	"jcc-golang-2021/model/shop"
	"jcc-golang-2021/model/user"
	"jcc-golang-2021/util"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func PostProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// parse form
	var item product.Product
	item.Name = r.PostFormValue("name")
	item.Detail = r.PostFormValue("detail")
	item.Category = r.PostFormValue("category")
	item.ImageUrl = r.PostFormValue("image_url")

	formShopId, err := strconv.Atoi(r.PostFormValue("shop_id"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	item.ShopId = formShopId

	formPrice, err := strconv.ParseFloat(r.PostFormValue("price"), 64)
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	item.Price = formPrice

	formStock, err := strconv.Atoi(r.PostFormValue("stock"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	item.Stock = formStock

	// check auth
	if !isProductBasicAuthValid(2, item, r, ctx) {
		err := map[string]string{
			"status": "Unauthorized!",
		}
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// post data
	if err := product.PostDb(ctx, item); err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Success!",
	}
	util.ResponseJSON(w, res, http.StatusOK)
}

func PutProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get item from db
	itemId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	theItem, err := product.GetByIdDb(ctx, itemId)
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
	formName := r.PostFormValue("name")
	formDetail := r.PostFormValue("detail")
	formCategory := r.PostFormValue("category")
	formImageUrl := r.PostFormValue("image_url")

	formShopId := r.PostFormValue("shop_id")
	formShopIdInt, err := strconv.Atoi(formShopId)
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	formPrice := r.PostFormValue("price")
	formPriceInt, err := strconv.ParseFloat(r.PostFormValue("price"), 64)
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	formStock := r.PostFormValue("stock")
	formStockInt, err := strconv.Atoi(formStock)
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	// update value
	if formName != "" {
		theItem[0].Name = formName
	}
	if formDetail != "" {
		theItem[0].Detail = formDetail
	}
	if formCategory != "" {
		theItem[0].Category = formCategory
	}
	if formImageUrl != "" {
		theItem[0].ImageUrl = formImageUrl
	}
	if formPrice != "" {
		theItem[0].Price = formPriceInt
	}
	if formStock != "" {
		theItem[0].Stock = formStockInt
	}
	if formShopId != "" {
		theItem[0].ShopId = formShopIdInt
	}

	// check auth
	if !isProductBasicAuthValid(2, theItem[0], r, ctx) {
		err := map[string]string{
			"status": "Unauthorized!",
		}
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// put db
	if err := product.PutDb(ctx, theItem[0]); err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Success!",
	}
	util.ResponseJSON(w, res, http.StatusOK)
}

func GetProductAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get param
	shopIdParam := r.URL.Query().Get("shop_id")
	shopIdInt, _ := strconv.Atoi(shopIdParam)

	// query
	var items []product.Product
	var err error
	if shopIdInt != 0 {
		items, err = product.GetByShopIdDb(ctx, shopIdInt)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		items, err = product.GetByIdDb(ctx, -1)
		if err != nil {
			fmt.Println(err)
		}
	}

	// result
	util.ResponseJSON(w, items, http.StatusOK)
}

func GetProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get id
	itemId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		itemId = -1
	}

	// query
	items, err := product.GetByIdDb(ctx, itemId)
	if err != nil {
		fmt.Println(err)
	}

	// result
	util.ResponseJSON(w, items, http.StatusOK)
}

func isProductBasicAuthValid(rule int, item product.Product, r *http.Request, ctx context.Context) bool {
	isValid := false
	theShop, err := shop.GetByIdDb(ctx, item.ShopId)
	if theShop != nil && err == nil {
		if user.IsBasicAuthValid(2, theShop[0].UserId, r, ctx) {
			isValid = true
		}
	}

	return isValid
}
