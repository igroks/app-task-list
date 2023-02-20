package models

type MatchRequest struct {
	Item *string `json:"item" binding:"required"`
}

type MatchResponse struct {
	Message string `json:"msg"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}
