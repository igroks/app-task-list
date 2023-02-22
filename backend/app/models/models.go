package models

import "time"

type Item struct {
	Id         int       `json:"id"`
	Name       string    `json:"name" binding:"required"`
	CreatedAt  time.Time `json:"createdAt" binding:"required"`
	Duplicated bool      `json:"duplicated"`
}

type DeleteItemsRequest struct {
	Id int `json:"id" binding:"required"`
}

type Response struct {
	Message string `json:"msg"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}
