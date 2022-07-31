package structs

type Response struct {
	Message string      `json:"message" form:"message"`
	Data    interface{} `json:"data"`
}
