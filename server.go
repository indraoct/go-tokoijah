package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/indraoct/go-tokoijah/apps"
	"time"
	"github.com/BurntSushi/toml"
	"fmt"
)


func initDb() *sql.DB {

	var conf apps.Config
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v\n", conf)

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.Database.User, conf.Database.Password, conf.Database.Server, conf.Database.Port, conf.Database.Database)

	db, err := sql.Open("mysql", connString)
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
	router.HandleFunc("/getproducts",apps.GetProducts(db)).Methods("GET")
	router.HandleFunc("/getproduct/{sku:[A-Za-z0-9-_/.]+}",apps.GetProduct(db)).Methods("GET")
	http.Handle("/",router)
	log.Fatal(http.ListenAndServe(":7777",router))



}

