package apps

import (
	"net/http"
	"database/sql"
	"encoding/json"
)

func GetProducthandler(db *sql.DB) http.HandlerFunc{
	function := func(writer http.ResponseWriter, request *http.Request){

		var responseProduct  ResponseProduct
		filter_sku := request.FormValue("sku")

		arr_products := GetProducts(db,filter_sku)
		responseProduct.Data = arr_products
		responseProduct.Status = 1
		responseProduct.Message = "Success"

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		json.NewEncoder(writer).Encode(responseProduct)

	}

	return http.HandlerFunc(function)
}
