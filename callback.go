package goapplib

import (
	"log"
	"context"
)

type Callback interface {
	OnAppCall(ctx context.Context, request *GoRequest) (response *GoResponse)
	OnAppReady()
	OnAppClose()
}

type defaultCallback struct{}

func (c *defaultCallback) OnAppReady() {
	log.Printf("[WARN] OnAppReady: not implemented")
}

func (c *defaultCallback) OnAppCall(ctx context.Context, request *GoRequest) (response *GoResponse) {
	log.Printf("[WARN] OnAppCall: %v, but not implemented", request)
	response = &GoResponse{
		TraceId: request.TraceId,
		Code:    Code_MethodNotFound,
		Data:    []byte("method not found"),
	}
	return
}

func (c *defaultCallback) OnAppClose() {
	log.Printf("[WARN] OnAppClose: not implemented")
}