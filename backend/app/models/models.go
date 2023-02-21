package models

type Request struct {
	Item *string `json:"item" binding:"required"`
}

type Response struct {
	Message string `json:"msg"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}
