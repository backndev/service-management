package dto

type Result struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Error      interface{} `json:"error"`
	Data       interface{} `json:"data"`
}
