package main

import "C"
import (
	"github.com/freedim-org/goapplib"
)

//export Address
func Address() *C.char {
	return C.CString(goapplib.Address())
}

func main() {}
