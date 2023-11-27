package goapplib

import "log"

type Callback interface {
	OnAppCall(request *Request) (response *Response)
}

type defaultCallback struct{}

func (c *defaultCallback) OnAppCall(request *Request) (response *Response) {
	log.Printf("[WARN] OnAppCall: %v, but not implemented", request)
	response = &Response{
		TraceId: request.TraceId,
		Code:    CodeMethodNotFound,
		Data:    "",
	}
	return
}
