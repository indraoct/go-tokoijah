package apps

import (
	"database/sql"
	"net/http"
	"github.com/heroku/go-getting-started/apps"
	"log"
	"encoding/json"
)

func ProductHandler(db *sql.DB) http.HandlerFunc{
	function := func(writer http.ResponseWriter, request *http.Request){

		var products apps.Products
		var arr_products []apps.Products
		var responseProduct apps.ResponseProduct

		rows, err := db.Query("SELECT sku,product_name,stocks from products ORDER BY sku DESC")
		if(err != nil){
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next(){
			if err := rows.Scan(&products.Sku, &products.Product_name, &products.Stocks); err != nil {
				log.Println(err.Error())
			}
			arr_products = append(arr_products,products)
		}

		responseProduct.Data = arr_products
		responseProduct.Status = 1
		responseProduct.Message = "Success"


		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		//writer.Write(byteInitial.Bytes())

		json.NewEncoder(writer).Encode(responseProduct)

	}

	return http.HandlerFunc(function)
}