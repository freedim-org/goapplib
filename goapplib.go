package goapplib

var server = NewLocalServer(&LocalServerConfig{
	Port: FreePort(),
})

func init() {
	server.Start()
}

func Port() int {
	return server.Config.Port
}
