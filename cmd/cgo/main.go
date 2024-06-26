package main

/*
This is a example of using goapplib in cgo.
*/

import "C"
import (
	"github.com/freedim-org/goapplib"
)

// init
// In your application, you must call goapplib.Init(xxx) before using goapplib.
func init() {
	goapplib.Init(&goapplib.LocalServerConfig{
		Address:  "",
		Callback: new(callback),
	})
}

//export Address
func Address() *C.char {
	return C.CString(goapplib.Address())
}

func main() {}

type callback struct{}

func (c *callback) OnAppReady() {}

func (c *callback) OnAppCall(request *goapplib.GoRequest) (response *goapplib.GoResponse) {
	response = &goapplib.GoResponse{
		TraceId: request.TraceId,
		Code:    goapplib.Code_OK,
		Data:    []byte{},
	}
	return
}
