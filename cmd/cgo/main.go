package main

import "C"
import "github.com/freedim-org/goapplib"

var server = goapplib.NewLocalServer(&goapplib.LocalServerConfig{
	Port: goapplib.FreePort(),
})

func init() {
	server.Start()
}

func Port() C.int {
	return C.int(server.Config.Port)
}

func main() {}
