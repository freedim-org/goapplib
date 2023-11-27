package goapplib

import "C"

var server = newLocalServer(&localServerConfig{
	Port: freePort(),
})

func init() {
	server.start()
}

func Port() int {
	return server.config.Port
}

func main() {}
