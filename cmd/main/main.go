package main

import (
	"fmt"
	"go/bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is up and running on 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}