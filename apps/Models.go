package apps

import (
	"database/sql"
	"net/http"
	"log"
	"encoding/json"
)

func GetProductHandler(db *sql.DB) http.HandlerFunc{
	function := func(writer http.ResponseWriter, request *http.Request){

		var products Products
		var arr_products [] Products
		var responseProduct  ResponseProduct
		filter_sku := request.FormValue("sku")

		if(filter_sku != ""){
			db.QueryRow("SELECT sku,product_name,stocks from products WHERE sku=?",filter_sku).Scan(&products.Sku,&products.Product_name,&products.Stocks)
                        arr_products = append(arr_products,products)

		}else {

			rows, err := db.Query("SELECT sku,product_name,stocks from products ORDER BY sku DESC")
			if (err != nil) {
				log.Fatal(err)
			}
			defer rows.Close()

			for rows.Next() {
				if err := rows.Scan(&products.Sku, &products.Product_name, &products.Stocks); err != nil {
					log.Println(err.Error())
				}
				arr_products = append(arr_products, products)
			}
		}

		responseProduct.Data = arr_products
		responseProduct.Status = 1
		responseProduct.Message = "Success"


		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		json.NewEncoder(writer).Encode(responseProduct)

	}

	return http.HandlerFunc(function)
}