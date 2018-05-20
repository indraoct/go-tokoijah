package main

import (
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/indraoct/go-tokoijah/apps"
)


func initDb() *sql.DB {

	db, err := sql.Open("mysql", "root:blink182@/tokoijah")
	if(err != nil){
		log.Fatal(err)
	}
	//set max idle conns
	db.SetMaxIdleConns(20)
	//setmaxopenconns
	db.SetMaxOpenConns(20)

	return db

}

/**
 * Main Function
 */
func main(){

	/**
	 * Initiate DB
	 */
	db := initDb()

	/**
	 * Router
	 */
	router := mux.NewRouter()
	router.HandleFunc("/getproducts",apps.ProductHandler(db)).Methods("GET")
	http.Handle("/",router)
	log.Fatal(http.ListenAndServe(":8000",router))



}

