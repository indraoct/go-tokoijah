package apps

import (
	"database/sql"
	"log"
)

func ProductsModel(db *sql.DB,filter map[string]string) []Products{
		var products Products
		var arr_products [] Products

		if(len(filter) > 0){
			err := db.QueryRow("SELECT sku,product_name,stocks from products WHERE sku=?",filter["sku"]).Scan(&products.Sku,&products.Product_name,&products.Stocks)
			if err != nil {
				log.Println(err)
			}

			arr_products = append(arr_products,products)
		}else {

			rows, err := db.Query("SELECT sku,product_name,stocks from products ORDER BY sku DESC")
			if err!= nil {
				log.Print(err)
			}
			defer rows.Close()

			for rows.Next() {
				if err := rows.Scan(&products.Sku, &products.Product_name, &products.Stocks); err != nil {
					log.Println(err.Error())
				}
				arr_products = append(arr_products, products)
			}
		}

	return arr_products

}