package route

import (
	"context"
	"fmt"
	"jcc-golang-2021/model/session"
	"jcc-golang-2021/model/user"
	"jcc-golang-2021/util"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get id
	itemId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		itemId = -1
	}

	// check auth
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

	// query
	items, err := user.GetByIdDb(ctx, itemId)
	if err != nil {
		fmt.Println(err)
	}
	util.ResponseJSON(w, items, http.StatusOK)
}

func GetUserAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// check auth
	if !user.IsBasicAuthValid(1, 0, r, ctx) {
		err := map[string]string{
			"status": "Unauthorized!",
		}
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// query
	items, err := user.GetByIdDb(ctx, -1)
	if err != nil {
		fmt.Println(err)
	}
	util.ResponseJSON(w, items, http.StatusOK)
}

func DelUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get user id
	userId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	// check auth
	if !user.IsBasicAuthValid(2, userId, r, ctx) {
		err := map[string]string{
			"status": "Unauthorized!",
		}
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// del user
	if err := user.DeactivateDb(ctx, userId); err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	// result
	res := map[string]string{
		"status": "Success!",
	}
	util.ResponseJSON(w, res, http.StatusOK)
}

func PutUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get item from db
	itemId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	theUser, err := user.GetByIdDb(ctx, itemId)
	if theUser == nil {
		err := map[string]string{
			"status": "User not found!",
		}
		util.ResponseJSON(w, err, http.StatusNotFound)
		return
	}
	if err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	// parse form
	FormFullName := r.PostFormValue("full_name")
	FormBirthDate := r.PostFormValue("birth_date")
	FormImageUrl := r.PostFormValue("image_url")
	FormPhone := r.PostFormValue("phone")
	FormEmail := r.PostFormValue("email")
	FormAddress := r.PostFormValue("address")

	// update value
	if FormFullName != "" {
		theUser[0].FullName = FormFullName
	}
	if FormBirthDate != "" {
		theUser[0].BirthDate = FormBirthDate
	}
	if FormImageUrl != "" {
		theUser[0].ImageUrl = FormImageUrl
	}
	if FormPhone != "" {
		theUser[0].Phone = FormPhone
	}
	if FormEmail != "" {
		theUser[0].Email = FormEmail
	}
	if FormAddress != "" {
		theUser[0].Address = FormAddress
	}

	// check auth
	if !user.IsBasicAuthValid(2, theUser[0].ID, r, ctx) {
		err := map[string]string{
			"status": "Unauthorized!",
		}
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// put db
	if err := user.PutDb(ctx, theUser[0]); err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Success!",
	}
	util.ResponseJSON(w, res, http.StatusOK)
}

func ResetPass(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// parse form
	FormOldPass := r.PostFormValue("old_pass")
	FormNewPass := r.PostFormValue("new_pass")
	FormConPass := r.PostFormValue("confirm_pass")

	// get item from db
	itemId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	theUser, err := user.GetByIdDb(ctx, itemId)
	if theUser == nil {
		err := map[string]string{
			"status": "User not found or password is invalid!",
		}
		util.ResponseJSON(w, err, http.StatusNotFound)
		return
	}
	if err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	// check confirm
	if FormNewPass != FormConPass {
		err := map[string]string{
			"status": "Confirm pass is invalid!",
		}
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	// check old pass
	passDecrypt := bcrypt.CompareHashAndPassword(
		[]byte(theUser[0].Password),
		[]byte(FormOldPass))
	if passDecrypt != nil {
		err := map[string]string{
			"status": "User not found or password is invalid!",
		}
		util.ResponseJSON(w, err, http.StatusNotFound)
		return
	}

	// set new pass
	newPassCrypted, err := bcrypt.GenerateFromPassword([]byte(FormNewPass), bcrypt.DefaultCost)
	if err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	theUser[0].Password = string(newPassCrypted)

	// check auth
	if !user.IsBasicAuthValid(2, theUser[0].ID, r, ctx) {
		err := map[string]string{
			"status": "Unauthorized!",
		}
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// put data
	if err := user.ResetPassDb(ctx, theUser[0]); err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Success!",
	}
	util.ResponseJSON(w, res, http.StatusOK)
}

func PostUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// parse form
	var item user.User
	item.Username = r.PostFormValue("username")
	item.Password = r.PostFormValue("password")
	item.FullName = r.PostFormValue("full_name")
	item.BirthDate = r.PostFormValue("birth_date")
	item.ImageUrl = r.PostFormValue("image_url")
	item.Phone = r.PostFormValue("phone")
	item.Email = r.PostFormValue("email")
	item.Address = r.PostFormValue("address")

	formRoleId, err := strconv.Atoi(r.PostFormValue("role_id"))
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	item.RoleId = formRoleId

	// check username not used
	isUserExisted, err := user.GetByUsernameDb(ctx, item.Username)
	if isUserExisted != nil {
		err := map[string]string{
			"status": "Username is already used!",
		}
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	if err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	// mod
	hash, err := bcrypt.GenerateFromPassword([]byte(item.Password), bcrypt.DefaultCost)
	if err != nil {
		util.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	item.Password = string(hash)

	// check auth
	if item.RoleId == 1 && !user.IsBasicAuthValid(1, 0, r, ctx) {
		err := map[string]string{
			"status": "You need to be logged in as Admin!",
		}
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// post data
	if err := user.PostDb(ctx, item); err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Success!",
	}
	util.ResponseJSON(w, res, http.StatusOK)

	// create new session
	theUser, err := user.GetByUsernameDb(ctx, item.Username)
	if theUser == nil {
		util.ResponseJSON(w, err, http.StatusNotFound)
		return
	}
	if err != nil {
		util.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	newSession := session.Session{
		ID:           0,
		Courier:      "",
		Note:         "",
		DeliveryCost: 0,
		UserId:       theUser[0].ID,
		StatusId:     0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now()}
	session.PostDb(ctx, newSession)
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// init
	var code int
	var res map[string]string
	var form user.User
	var isValid = false

	// defer
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// parse body
	form.Username = r.PostFormValue("username")
	form.Password = r.PostFormValue("password")

	// check data
	theUser, err := user.GetByUsernameDb(ctx, form.Username)
	if err == nil && theUser != nil {
		passDecrypt := bcrypt.CompareHashAndPassword(
			[]byte(theUser[0].Password),
			[]byte(form.Password))

		if passDecrypt == nil {
			isValid = true
		}
	}

	if isValid {
		code = http.StatusOK
		res = map[string]string{
			"status": "Success!",
		}
	} else {
		code = http.StatusNotFound
		res = map[string]string{
			"status": "Username or Password is invalid!",
		}
	}

	// result
	util.ResponseJSON(w, res, code)
}
