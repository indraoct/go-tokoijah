package main

import (
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	"net/http"
	"github.com/indraoct/go-tokoijah/apps"
)



/**
 * Main Function
 */
func main(){

	router := mux.NewRouter()

	router.HandleFunc("/getproducts",apps.GetProducts)

	log.Fatal(http.ListenAndServe(":9090",router))


}
