package success

type ApiResponse struct {
	Item   interface{} `json:"item"`
	Status int         `json:"status"`
}
