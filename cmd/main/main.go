package main

import (
	"log"
	"net/http"

	_ "githb.com/jinzhu/gorm/dialect/mysql"
	"github.com/gorilla/mux"
	"github.com/imran1509/book-management-system/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
