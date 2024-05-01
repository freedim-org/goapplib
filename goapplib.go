package goapplib

var server *LocalServer

func Init(c *LocalServerConfig) {
	dft := DefaultLocalServerConfig()
	if c == nil {
		c = dft
	}
	if c.Address == "" {
		c.Address = dft.Address
	}
	if c.Callback == nil {
		c.Callback = dft.Callback
	}
	server = NewLocalServer(c)
	server.Start()
}

func CallApp(req *GoRequest) *GoResponse {
	return server.callApp(req)
}

func Address() string {
	return server.Config.Address
}
