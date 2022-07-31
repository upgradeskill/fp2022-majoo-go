package structs

type Response struct {
	Message string      `json:"message" form:"message"`
	Data    interface{} `json:"data"`
}

type ResponsePagination struct {
	Message string      `json:"message" form:"message"`
	Data    interface{} `json:"data"`
	Limit   int         `json:"limit"`
	Offset  int         `json:"offset"`
}
