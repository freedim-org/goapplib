package tools

import "net"

func FreePort() int32 {
	// 获取空闲端口
	var a *net.TCPAddr
	var err error
	if a, err = net.ResolveTCPAddr("tcp", "localhost:0"); err == nil {
		var l *net.TCPListener
		if l, err = net.ListenTCP("tcp", a); err == nil {
			defer l.Close()
			return int32(l.Addr().(*net.TCPAddr).Port)
		}
	}
	panic(err)
}
