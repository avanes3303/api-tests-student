package models

// ErrorResponse описывает код и сообщение ошибки API.
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
