package goapplib

import (
	"fmt"
	"github.com/freedim-org/goapplib/dp"
	"github.com/freedim-org/goapplib/tools"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

/*
Path: localserver.go
Start a local TCP server for communication with the caller.
*/

type LocalServerConfig struct {
	Address  string
	Callback Callback
}

type LocalServer struct {
	client          *net.TCPConn
	listener        *net.TCPListener
	Config          *LocalServerConfig
	responseChanMap sync.Map
}

func DefaultLocalServerConfig() *LocalServerConfig {
	return &LocalServerConfig{
		Address:  fmt.Sprintf("127.0.0.1:%d", tools.FreePort()),
		Callback: new(defaultCallback),
	}
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
		if msg.IsResp() {
			// response
			l.onNewResponse(msg.GetData())
		} else {
			l.onNewRequest(msg.GetData())
		}
	}
}

func (l *LocalServer) onNewRequest(data string) {
	request := &Request{}
	err := request.Unmarshal(data)
	if err != nil {
		log.Printf("[ERRO] request.Unmarshal: %v", err)
	} else {
		response := l.Config.Callback.OnAppCall(request)
		if response == nil {
			response = &Response{
				TraceId: request.TraceId,
				Code:    CodeMethodNullResponse,
				Data:    "",
			}
		}
		response.TraceId = request.TraceId
		// send response
		l.sendResponse(response)
	}
}

func (l *LocalServer) callApp(req *Request) *Response {
	data := req.Marshal()
	pack, err := dp.DP.Pack(&dp.Message{
		Len:        uint32(len(data)),
		IsResponse: false,
		Data:       data,
	})
	if err != nil {
		panic(err)
	}
	ch := make(chan *Response)
	l.responseChanMap.Store(req.TraceId, ch)
	defer l.responseChanMap.Delete(req.TraceId)
	_, err = l.client.Write(pack)
	if err != nil {
		panic(err)
	}
	response := <-ch
	return response
}

func (l *LocalServer) onNewResponse(data string) {
	response := &Response{}
	err := response.Unmarshal(data)
	if err != nil {
		log.Printf("[ERRO] response.Unmarshal: %v", err)
	} else {
		ch, ok := l.responseChanMap.Load(response.TraceId)
		if !ok {
			log.Printf("[WARN] responseChanMap.Load: not found")
			return
		}
		ch.(chan *Response) <- response
	}
}

func (l *LocalServer) sendResponse(response *Response) {
	data := response.Marshal()
	pack, err := dp.DP.Pack(&dp.Message{
		Len:        uint32(len(data)),
		IsResponse: true,
		Data:       data,
	})
	if err != nil {
		log.Printf("[ERRO] dp.DP.Pack: %v", err)
		return
	}
	_, err = l.client.Write(pack)
	if err != nil {
		log.Printf("[ERRO] client.Write: %v", err)
		return
	}
}
