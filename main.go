package main

import (
	"fmt"
	"jcc-golang-2021/config"
	"jcc-golang-2021/route"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

var (
// testUsers      = []model.User{}
// testPersons    = []model.Person{}
// testRoles      = []role.Role{}
// testProducts   = []model.Product{}
// testDevelopers = []model.Developer{}
// testLibraries  = []model.Library{}
// testReviews    = []model.Review{}
)

func main() {
	// db check
	db, e := config.MySQL()
	if e != nil {
		log.Fatal(e)
	}
	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}
	fmt.Println("DB connected!")

	// port check
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// routes
	router := httprouter.New()
	router.NotFound = http.FileServer(http.Dir("Public"))

	// user
	router.POST("/api/user/register", route.PostUserRegister)
	router.GET("/api/user/:id", route.GetUser)
	router.PUT("/api/user/:id", route.PutUser)
	router.PUT("/api/user/:id/reset-pass", route.ResetPass)
	router.DELETE("/api/user/:id", route.DelUser)

	// role
	router.GET("/api/role", route.GetRole)

	// shop
	router.POST("/api/shop/register", route.PostShopRegister)
	router.GET("/api/shop/:id", route.GetShop)
	router.PUT("/api/shop/:id", route.PutShop)
	router.PUT("/api/shop/:id/toggle-active", route.PutShopToggleActive)

	// product
	router.POST("/api/product", route.PostProduct)
	router.GET("/api/product/:id", route.GetProduct)
	router.PUT("/api/product/:id", route.PutProduct)

	// review
	router.POST("/api/review", route.PostReview)
	router.GET("/api/review/:id", route.GetReview)
	router.PUT("/api/review/:id", route.PutReview)
	router.PUT("/api/review/:id/response", route.PutResponse)

	// session
	router.GET("/api/session/:id", route.GetSession)
	router.PUT("/api/session/:id", route.PutSession)

	// status
	router.GET("/api/status", route.GetStatus)

	// cart
	router.POST("/api/cart", route.PostCart)
	router.GET("/api/cart/:id", route.GetCart)
	router.PUT("/api/cart/:id", route.PutCart)
	router.DELETE("/api/cart/:id", route.DelCart)

	// not used
	/*
		router.POST("/api/session", route.PostSession)
		router.POST("/api/login", route.Login)
	*/

	// serve
	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
