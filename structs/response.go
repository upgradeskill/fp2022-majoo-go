package structs

type Response struct {
	Status  int         `json:"status" form:"status"`
	Message string      `json:"message" form:"message"`
	Data    interface{} `json:"data"`
}
