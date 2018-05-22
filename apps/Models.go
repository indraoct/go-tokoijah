package apps

import (
	"database/sql"
	"log"
)

func ProductsModel(db *sql.DB,filter map[string]string) ReturnDataProduct{
		var products Products
		var arr_products [] Products
		var returnDataProduct ReturnDataProduct

		if(len(filter) > 0){
			err := db.QueryRow("SELECT sku,product_name,stocks from products WHERE sku=?",filter["sku"]).Scan(&products.Sku,&products.Product_name,&products.Stocks)
			arr_products = append(arr_products,products)
			if err != nil {
				log.Println(err)
				returnDataProduct.Count = 0
			}else{
				returnDataProduct.Count = len(arr_products)
			}
			returnDataProduct.Data = arr_products

		}else {

			rows, err := db.Query("SELECT sku,product_name,stocks from products ORDER BY sku DESC")
			if err!= nil {
				log.Print(err)
			}
			defer rows.Close()

			count := 0
			for rows.Next() {
				if err := rows.Scan(&products.Sku, &products.Product_name, &products.Stocks); err != nil {
					log.Println(err.Error())

				}else{
					arr_products = append(arr_products, products)
					count++
				}

			}

			returnDataProduct.Data = arr_products
			returnDataProduct.Count = count
		}

	return returnDataProduct

}