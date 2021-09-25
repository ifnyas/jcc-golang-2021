package route

import (
	"context"
	"fmt"
	"jcc-golang-2021/model/role"
	"jcc-golang-2021/model/user"
	"jcc-golang-2021/util"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetRole(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get id
	itemId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		itemId = -1
	}

	// check auth
	if !user.IsBasicAuthValid(1, 0, r, ctx) {
		err := map[string]string{
			"status": "Unauthorized!",
		}
		util.ResponseJSON(w, err, http.StatusUnauthorized)
		return
	}

	// query
	items, err := role.GetByIdDb(ctx, itemId)
	if err != nil {
		fmt.Println(err)
	}
	util.ResponseJSON(w, items, http.StatusOK)
}
