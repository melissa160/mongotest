package models

// Response message
type Response struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Error code
type ErrorCode struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}
