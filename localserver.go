package goapplib

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

/*
Path: localserver.go
Start a local TCP server for communication with the caller.
*/

type LocalServerConfig struct {
	Port int
}

type localServer struct {
	client   *net.TCPConn
	listener *net.TCPListener
	Config   *LocalServerConfig
}

func NewLocalServer(config *LocalServerConfig) *localServer {
	l := &localServer{
		Config: config,
	}
	return l
}

func (l *localServer) Start() {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%d", l.Config.Port))
	if err != nil {
		panic(err)
	}
	l.listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	go func() {
		err = l.acceptOnce()
		if err != nil {
			panic(err)
		}
		l.loopRead()
	}()
	return
}

func (l *localServer) acceptOnce() error {
	conn, err := l.listener.AcceptTCP()
	if err != nil {
		return err
	}
	l.client = conn
	return nil
}

func (l *localServer) loopRead() {
	for {
		msg, err := DP.Unpack(l.client)
		if err != nil {
			if err == io.EOF {
				// connection closed
				os.Exit(0)
			}
			// use of closed network connection
			if strings.HasSuffix(err.Error(), "use of closed network connection") {
				os.Exit(0)
			}
			panic(err)
		}
		l.onNewData(msg.GetData())
	}
}

func (l *localServer) onNewData(data string) {

}
