package goapplib

import "encoding/json"

type Request struct {
	TraceId string `json:"traceId"`
	Method  string `json:"method"`
	Data    string `json:"data"`
}

func (r *Request) Unmarshal(data string) error {
	return json.Unmarshal([]byte(data), r)
}

func (r *Request) Marshal() string {
	data, _ := json.Marshal(r)
	return string(data)
}

func (r *Request) Parse(in any) error {
	return json.Unmarshal([]byte(r.Data), in)
}

func (r *Request) ParseError(err error) *Response {
	return &Response{
		TraceId: r.TraceId,
		Code:    CodeInvalidRequest,
		Data:    err.Error(),
	}
}

type Code int

const (
	CodeOK                 Code = 0
	CodeInvalidRequest     Code = 400
	CodeMethodNotFound     Code = 404
	CodeInternalError      Code = 500
	CodeMethodNullResponse Code = 501
)

type Response struct {
	TraceId string `json:"traceId"`
	Code    Code   `json:"code"`
	Data    string `json:"data"`
}

func (r *Response) Marshal() string {
	data, _ := json.Marshal(r)
	return string(data)
}

func (r *Response) Unmarshal(data string) error {
	return json.Unmarshal([]byte(data), r)
}
