package main

import (
	"fmt"
	"jcc-golang-2021/route"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// routes
	router := httprouter.New()
	router.NotFound = http.RedirectHandler("/static", http.StatusMovedPermanently)
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	router.GET("/search", route.Search)

	// port check
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// serve
	fmt.Println("Server Running at Port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
