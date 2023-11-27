package goapplib

import "fmt"

var server = NewLocalServer(&LocalServerConfig{
	Address: fmt.Sprintf("127.0.0.1:%d", FreePort()),
})

func init() {
	server.Start()
}

func Address() string {
	return server.Config.Address
}
