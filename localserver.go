package goapplib

import (
	"github.com/freedim-org/goapplib/dp"
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
	Address string
}

type LocalServer struct {
	client   *net.TCPConn
	listener *net.TCPListener
	Config   *LocalServerConfig
}

func NewLocalServer(config *LocalServerConfig) *LocalServer {
	l := &LocalServer{
		Config: config,
	}
	return l
}

func (l *LocalServer) Start() {
	addr, err := net.ResolveTCPAddr("tcp", l.Config.Address)
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

func (l *LocalServer) acceptOnce() error {
	conn, err := l.listener.AcceptTCP()
	if err != nil {
		return err
	}
	l.client = conn
	return nil
}

func (l *LocalServer) loopRead() {
	for {
		msg, err := dp.DP.Unpack(l.client)
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

func (l *LocalServer) onNewData(data string) {

}
