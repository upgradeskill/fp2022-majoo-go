package structs

type TransactionDetails struct {
	Id            int      `json:"id" form:"id" gorm:"primaryKey"`
	TransactionId int      `json:"transaction_id" form:"transaction_id"`
	ProductId     int      `json:"product_id" form:"product_id"`
	Quantity      int      `json:"quantity" form:"quantity"`
	Price         int      `json:"price" form:"price"`
	Note          string   `json:"note" form:"note"`
	Product       Products `gorm:"Foreignkey:product_id;" json:"product"`
}
