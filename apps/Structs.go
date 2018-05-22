package apps


/**
 * Group Of struct
 */

//Config Struct
type Config struct {
	Output Output
	Database Database
}

type Database struct {
	Server string
	Port string
	Database string
	User string
	Password string
}

type Output struct {
	Directory string
	Format string
}

//Products struct
type Products struct {
	Sku          string    `gorm:"not null" form:"sku" json:"sku"`
	Product_name string    `gorm:"not null" form:"product_name" json:"product_name"`
	Stocks       int    `gorm:"not null" form:"stocks" json:"stocks"`
}

type ReturnDataProduct struct {
	Data         []Products `form:"data" json:"data"`
	Count        int        `form:"count" json:"count"`

}


//Stock_Ins
type Stock_Ins struct {
	Sku               string    `gorm:"not null" form:"sku" json:"sku"`
	Created_Date      string    `gorm:"not null" form:"created_date" json:"created_date"`
	Buy_Price         int       `gorm:"not null" form:"buy_price" json:"buy_price"`
	Qty               int       `gorm:"not null" form:"qty" json:"qty"`
	Kwitansi          string    `form:"kwitansi" json:"kwitansi"`

}


//Stock_Outs
type Stock_Outs struct {
	Transaction_Id string `gorm:"null" form"transaction_id" json:"transaction_id"`
	Sku               string    `gorm:"not null" form:"sku" json:"sku"`
	Qty               int       `gorm:"not null" form:"qty" json:"qty"`
	Note              string    `gorm:"not null" form:"note" json:"note"`
	Created_Date      string    `gorm:"not null" form:"created_date" json:"created_date"`

}

//Stock_Ins_CSV
type Stock_Ins_CSV struct {
	Sku               string    `csv:"sku"`
	Created_Date      string    `csv:"created_date"`
	Buy_Price         string    `csv:"buy_price"`
	Qty               string    `csv:"qty"`
	Kwitansi          string    `csv:"kwitansi"`

}

//Stock_Outs_CSV
type Stock_Outs_CSV struct {
	Transaction_Id string       `csv:"transaction_id"`
	Sku               string    `csv:"sku"`
	Qty               string    `csv:"qty"`
	Note              string    `csv:"note"`
	Created_Date      string    `csv:"created_date"`

}

//transactions
type Transactions struct {
	Id                string `gorm:"null" form"id" json:"id"`
	Sku               string    `gorm:"not null" form:"sku" json:"sku"`
	Qty               int       `gorm:"not null" form:"qty" json:"qty"`
	Buy_Price         int       `gorm:"not null" form:"buy_price" json:"buy_price"`
	Sell_Price        int       `gorm:"not null" form:"sell_price" json:"sell_price"`
	Created_Date      string    `gorm:"not null" form:"created_date" json:"created_date"`
}

//Transaction CSV

type Transactions_CSV struct {
	Id                string             `csv:"id"`
	Created_Date      string             `csv:"created_date"`
	Sku               string             `csv:"sku"`
	Product_name      string             `csv:"product_name"`
	Qty               string             `csv:"qty"`
	Buy_Price         string             `csv:"buy_price"`
	Total             string             `csv:"total"`
	Sell_Price        string             `csv:"sell_price"`
	Laba              string             `csv:"laba"`



}

//Get Product valuation CSV

type ProductValuation_CSV struct {
	Sku    string `csv:"sku"`
	Product_name      string             `csv:"product_name"`
	Qty               string             `csv:"qty"`
	Avg_Buy_Price     string             `csv:"avg_buy_price"`
	Total             string             `csv:"total"`


}

//ResponseProduct Struct
type ResponseProduct struct {
	Status    int        `form:"status" json:"status"`
	Message   string     `form:"message" json:"message"`
	Data      []Products `form:"data" json:"data"`
	Count     int        `form:"count" json:"count"`
}

//Response Insert Product
type ResponseInsertProduct struct {
	Status    int        `form:"status" json:"status"`
	Message   string     `form:"message" json:"message"`
	Data      DataInsertProduct
}

//data insert product
type DataInsertProduct struct{
	Sku               string    `gorm:"not null" form:"sku" json:"sku"`
	Product_name      string    `gorm:"not null" form:"product_name" json:"product_name"`
	Stocks            int       `gorm:"not null" form:"stocks" json:"stocks"`
	Buy_Price         int       `gorm:"not null" form:"buy_price" json:"buy_price"`
	Created_Date      string    `gorm:"not null" form:"created_date" json:"created_date"`
}

//ResponseTransaction
type ResponseTransaction struct {
	Status    int        `form:"status" json:"status"`
	Message   string     `form:"message" json:"message"`
	Data      DataTransaction

}

//Data Transaction
type DataTransaction struct {
	Transaction_Id    string    `form:"transaction_id" json:"transaction_id"`
	Sku               string    `gorm:"not null" form:"sku" json:"sku"`
	Product_name      string    `gorm:"not null" form:"product_name" json:"product_name"`
	Stocks            int       `gorm:"not null" form:"stocks" json:"stocks"`
	Buy_Price         int       `gorm:"not null" form:"buy_price" json:"buy_price"`
	Sell_Price        int       `gorm:"not null" form:"sell_price" json:"sell_price"`
	Created_Date      string    `gorm:"not null" form:"created_date" json:"created_date"`
}
