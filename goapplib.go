package goapplib

import "fmt"

var server = NewLocalServer(&LocalServerConfig{
	Port: FreePort(),
})

func init() {
	server.Start()
}

func Address() string {
	return fmt.Sprintf("127.0.0.1:%d", server.Config.Port)
}
