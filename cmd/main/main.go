package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rahul/go-bookstore/pkg/routes"
	// "github.com/rahul/go-bookstore/pkg/routes"

	_ "gorm.io/driver/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("Localhost: 9010", r))
	// ListenAndServe is a func which helps to create server in first place
}
