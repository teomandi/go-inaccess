package structs

type ErrorResponse struct {
	Status string `json:"status"`
	Desc   string `json:"desc"`
}
