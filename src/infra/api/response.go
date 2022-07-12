package api

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Status  int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
}
