package entity

type Response struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Error      string      `json:"error"`
	Data       interface{} `json:"data"`
}
