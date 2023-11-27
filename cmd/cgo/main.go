package main

import "C"
import (
	"github.com/freedim-org/goapplib"
)

var server = goapplib.NewLocalServer(&goapplib.LocalServerConfig{
	Port: goapplib.FreePort(),
})

func init() {
	server.Start()
}

//export Address
func Address() *C.char {
	return C.CString(goapplib.Address())
}

func main() {}
