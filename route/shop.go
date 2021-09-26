package route

import (
	"context"
	"fmt"
	"jcc-golang-2021/model/shop"
	"jcc-golang-2021/model/user"
	"jcc-golang-2021/util"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func PostShop(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// parse form
	var item shop.Shop
	item.Tag = r.PostFormValue("tag")
	item.Name = r.PostFormValue("name")
	item.Detail = r.PostFormValue("detail")
	item.ImageUrl = r.PostFormValue("image_url")
	item.Phone = r.PostFormValue("phone")
	item.Email = r.PostFormValue("email")
	item.Address = r.PostFormValue("address")
	item.IsActive = 1

	theUser := user.GetByBasicAuth(ctx, r)
	if theUser != nil {
		item.UserId = theUser[0].ID
	} else {
		err := map[string]string{
			"status": "Unauthorized!",
		}
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	formUserId := r.PostFormValue("user_id")
	formUserIdInt, err := strconv.Atoi(formUserId)
	if err == nil && theUser[0].RoleId == 1 {
		item.UserId = formUserIdInt
	}
	fmt.Println(item.UserId, formUserIdInt)
	if theUser[0].RoleId != 1 && item.UserId != formUserIdInt && formUserIdInt != 0 {
		err := map[string]string{
			"status": "Unauthorized!",
		}
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// check tag not used
	isShopExisted, err := shop.GetByTagDb(ctx, item.Tag)
	if isShopExisted != nil {
		err := map[string]string{
			"status": "Tag is already used!",
		}
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	if err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
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
	if err := shop.PostDb(ctx, item); err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Success!",
	}
	util.ResponseJSON(w, res, http.StatusOK)
}

func PutShopToggleActive(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get item from db
	itemId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	theItem, err := shop.GetByIdDb(ctx, itemId)
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

	// update value
	newValue := 0
	if theItem[0].IsActive == 0 {
		newValue = 1
	}
	theItem[0].IsActive = newValue

	// check auth
	if !user.IsBasicAuthValid(2, theItem[0].ID, r, ctx) {
		err := map[string]string{
			"status": "Unauthorized!",
		}
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// put db
	if err := shop.PutDb(ctx, theItem[0]); err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Success!",
	}
	util.ResponseJSON(w, res, http.StatusOK)
}

func PutShop(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get item from db
	itemId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	theItem, err := shop.GetByIdDb(ctx, itemId)
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
	formImageUrl := r.PostFormValue("image_url")
	formPhone := r.PostFormValue("phone")
	formEmail := r.PostFormValue("email")
	formAddress := r.PostFormValue("address")

	// update value
	if formName != "" {
		theItem[0].Name = formName
	}
	if formDetail != "" {
		theItem[0].Detail = formDetail
	}
	if formImageUrl != "" {
		theItem[0].ImageUrl = formImageUrl
	}
	if formPhone != "" {
		theItem[0].Phone = formPhone
	}
	if formEmail != "" {
		theItem[0].Email = formEmail
	}
	if formAddress != "" {
		theItem[0].Address = formAddress
	}

	// check auth
	if !user.IsBasicAuthValid(2, theItem[0].ID, r, ctx) {
		err := map[string]string{
			"status": "Unauthorized!",
		}
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// put db
	if err := shop.PutDb(ctx, theItem[0]); err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Success!",
	}
	util.ResponseJSON(w, res, http.StatusOK)
}

func GetShop(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get id
	itemId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	// check auth
	/*
		rule := 2
		if itemId < 0 {
			rule = 1
		}
		if !user.IsBasicAuthValid(rule, itemId, r, ctx) {
			err := map[string]string{
				"status": "Unauthorized!",
			}
			util.ResponseJSON(w, err, http.StatusUnauthorized)
			return
		}
	*/

	// result
	items, err := shop.GetByIdDb(ctx, itemId)
	if err != nil {
		fmt.Println(err)
	}
	util.ResponseJSON(w, items, http.StatusOK)
}

func GetShopAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// result
	items, err := shop.GetDb(ctx)
	if err != nil {
		fmt.Println(err)
	}
	util.ResponseJSON(w, items, http.StatusOK)
}
