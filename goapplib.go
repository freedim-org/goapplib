package goapplib

import (
	"fmt"
	"github.com/freedim-org/goapplib/tools"
)

var server = NewLocalServer(&LocalServerConfig{
	Address: fmt.Sprintf("127.0.0.1:%d", tools.FreePort()),
})

func init() {
	server.Start()
}

func Address() string {
	return server.Config.Address
}
