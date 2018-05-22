package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/indraoct/go-tokoijah/apps"
	"time"
)


func initDb() *sql.DB {

	db, err := sql.Open("mysql", "root:blink182@/tokoijah")
	if(err != nil){
		log.Fatal(err)
	}

	db.SetMaxIdleConns(0)

	db.SetMaxOpenConns(5)

	db.SetConnMaxLifetime(time.Minute*5)

	return db

}


func main(){

	db := initDb()

	router := mux.NewRouter()
	router.HandleFunc("/getproducts",apps.GetProductshandler(db)).Methods("GET")
	router.HandleFunc("/getproduct/{id}",apps.GetProducthandler(db)).Methods("GET")
	http.Handle("/",router)
	log.Fatal(http.ListenAndServe(":7777",router))



}

