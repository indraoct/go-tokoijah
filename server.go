package main

import (
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
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

	db.SetMaxIdleConns(1)

	db.SetMaxOpenConns(5)

	db.SetConnMaxLifetime(time.Minute*5)

	return db

}


func main(){

	db := initDb()

	router := mux.NewRouter()
	router.HandleFunc("/getproducts",apps.GetProductHandler(db)).Methods("GET")
	http.Handle("/",router)
	log.Fatal(http.ListenAndServe(":8000",router))



}

