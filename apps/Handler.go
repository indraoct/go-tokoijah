package apps

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
)

/**
 * Get rows data product
 */
func GetProductshandler(db *sql.DB) http.HandlerFunc{
	function := func(writer http.ResponseWriter, request *http.Request){

		var responseProduct  ResponseProduct
		filter := make(map[string]string)

		arr_products := GetProducts(db,filter)
		responseProduct.Data = arr_products
		responseProduct.Status = 1
		responseProduct.Message = "Success"

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		json.NewEncoder(writer).Encode(responseProduct)

	}

	return http.HandlerFunc(function)
}

/**
 * Single Row Product SKU
 */

func GetProducthandler(db *sql.DB) http.HandlerFunc{
	function := func(writer http.ResponseWriter,request *http.Request){

		var responseProduct ResponseProduct
		sku := mux.Vars(request)["sku"]
		filter := make(map[string]string)
		filter["sku"] = sku

		if(sku != ""){
			arr_products := GetProducts(db,filter)
			responseProduct.Status = 1
			responseProduct.Message = "Success"
			responseProduct.Data = arr_products

		}else{
			responseProduct.Status = 0
			responseProduct.Message = "SKU can not be empty"
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(responseProduct)

	}
	return http.HandlerFunc(function)

}
