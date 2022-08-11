package structs

type Transactions struct {
	Id                 int                  `json:"id" form:"id" gorm:"primaryKey"`
	OutletId           int                  `json:"outlet_id" form:"outlet_id"`
	Code               string               `json:"code" form:"code"`
	CustomerName       string               `json:"customer_name" form:"customer_name"`
	CreatedBy          int                  `json:"created_by" form:"created_by"`
	Outlet             Outlets              `gorm:"Foreignkey:outlet_id;" json:"outlet"`
	TransactionDetails []TransactionDetails `gorm:"Foreignkey:transaction_id;association_foreignkey:Id;" json:"transaction_details"`
}
