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
func GetProducts(db *sql.DB) http.HandlerFunc{
	function := func(writer http.ResponseWriter, request *http.Request){

		var responseProduct  ResponseProduct
		filter := make(map[string]string)

		rtn_products := ProductsModel(db,filter)
		responseProduct.Data = rtn_products.Data
		responseProduct.Status = 1
		responseProduct.Message = "Success"
		responseProduct.Count = rtn_products.Count

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		json.NewEncoder(writer).Encode(responseProduct)

	}

	return http.HandlerFunc(function)
}

/**
 * Single Row Product SKU
 */

func GetProduct(db *sql.DB) http.HandlerFunc{
	function := func(writer http.ResponseWriter,request *http.Request){

		var responseProduct ResponseProduct
		sku := mux.Vars(request)["sku"]
		filter := make(map[string]string)
		filter["sku"] = sku

		if(sku != ""){
			rtn_products := ProductsModel(db,filter)
			responseProduct.Status = 1
			responseProduct.Message = "Success"
			responseProduct.Data = rtn_products.Data
			responseProduct.Count = rtn_products.Count

		}else{
			responseProduct.Status = 0
			responseProduct.Message = "SKU can not be empty"
			responseProduct.Count = 0
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(responseProduct)

	}
	return http.HandlerFunc(function)

}
