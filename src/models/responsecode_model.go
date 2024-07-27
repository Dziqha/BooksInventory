package models

type ResponseCode struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}