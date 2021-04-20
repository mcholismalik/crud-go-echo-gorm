package model

import (
	"net/http"
	"time"
)

type LogError struct {
	Id            string      `json:"id"`
	RequestHeader http.Header `json:"request_header"`
	RequestBody   string      `json:"request_body"`
	URL           string      `json:"url"`
	HttpMethod    string      `json:"http_method"`
	FunctionName  string      `json:"function_name"`
	ErrorMessage  string      `json:"error_message"`
	Level         string      `json:"level"`
	CreatedAt     time.Time   `json:"created_at"`
}

type LogActivity struct {
	Id            string      `json:"id"`
	RequestHeader http.Header `json:"request_header"`
	RequestBody   string      `json:"request_body"`
	ResponseBody  string      `json:"response_body"`
	URL           string      `json:"url"`
	HttpMethod    string      `json:"http_method"`
	FunctionName  string      `json:"function_name"`
	CreatedAt     time.Time   `json:"created_at"`
}

type LogLogin struct {
	Id            string      `json:"id"`
	RequestHeader http.Header `json:"request_header"`
	RequestBody   string      `json:"request_body"`
	URL           string      `json:"url"`
	HttpMethod    string      `json:"http_method"`
	UserID        string      `json:"user_id"`
	CreatedAt     string      `json:"created_at"`
}
