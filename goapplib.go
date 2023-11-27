package goapplib

var server = NewLocalServer(&LocalServerConfig{
	Port: FreePort(),
})

func init() {
	server.Start()
}

func Port() int32 {
	return server.Config.Port
}
