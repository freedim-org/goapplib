package goapplib

import (
	"fmt"
	"github.com/freedim-org/goapplib/dp"
	"github.com/freedim-org/goapplib/tools"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"context"
	"time"
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
	requestCancelMap sync.Map
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
		for {
			err = l.acceptOnce()
			if err != nil {
				panic(err)
			}
			go l.Config.Callback.OnAppReady()
			l.loopRead()
			l.Config.Callback.OnAppClose()
		}
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
				return
			}
			// use of closed network connection
			if strings.HasSuffix(err.Error(), "use of closed network connection") {
				return
			}
			return
		}
		if msg.IsResp() {
			// response
			go l.onNewResponse(msg.GetData())
		} else {
			go l.onNewRequest(msg.GetData())
		}
	}
}

func (l *LocalServer) onNewRequest(data []byte) {
	request := &GoRequest{}
	err := request.Unmarshal(data)
	if err != nil {
		log.Printf("[ERRO] request.Unmarshal: %v", err)
	} else {
		if request.Method == "CancelToken" {
			cancelI, ok := l.requestCancelMap.Load(request.TraceId)
			if ok {
				cancel := cancelI.(func())
				cancel()
				l.requestCancelMap.Delete(request.TraceId)
				return
			}
		}
		ctx, cancel := context.WithCancel(context.Background())
		l.requestCancelMap.Store(request.TraceId, cancel)
		response := l.Config.Callback.OnAppCall(ctx, request)
		if response == nil {
			response = &GoResponse{
				TraceId: request.TraceId,
				Code:    Code_MethodNullResponse,
				Data:    []byte{},
			}
		}
		response.TraceId = request.TraceId
		// send response
		l.sendResponse(response)
	}
}

func (l *LocalServer) callApp(req *GoRequest) *GoResponse {
	data := req.Marshal()
	pack, err := dp.DP.Pack(&dp.Message{
		Len:        uint32(len(data)),
		IsResponse: false,
		Data:       data,
	})
	if err != nil {
		panic(err)
	}
	ch := make(chan *GoResponse)
	l.responseChanMap.Store(req.TraceId, ch)
	defer l.responseChanMap.Delete(req.TraceId)
	_, _ = l.client.Write(pack)
	select {
	case <-time.After(5 * time.Second):
		return &GoResponse{
			TraceId: req.TraceId,
			Code:    Code_Canceled,
		}
	case response := <-ch:
		return response
	}
}

func (l *LocalServer) onNewResponse(data []byte) {
	response := &GoResponse{}
	err := response.Unmarshal(data)
	if err != nil {
		log.Printf("[ERRO] response.Unmarshal: %v", err)
	} else {
		ch, ok := l.responseChanMap.Load(response.TraceId)
		if !ok {
			log.Printf("[WARN] responseChanMap.Load: not found")
			return
		}
		ch.(chan *GoResponse) <- response
	}
}

func (l *LocalServer) sendResponse(response *GoResponse) {
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
