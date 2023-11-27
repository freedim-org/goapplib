package main

import (
	"fmt"
	"github.com/freedim-org/goapplib"
	"github.com/freedim-org/goapplib/dp"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type callback struct{}

func (c *callback) OnAppCall(request *goapplib.Request) (response *goapplib.Response) {
	switch request.Method {
	case "test", "":
		response = &goapplib.Response{
			Code: goapplib.CodeOK,
			Data: request.Data,
		}
	default:
		response = &goapplib.Response{
			Code: goapplib.CodeMethodNotFound,
			Data: "",
		}
	}
	return
}

var traceId *atomic.Int64

func init() {
	traceId = new(atomic.Int64)
}

func NextTraceId() string {
	id := traceId.Add(1)
	return strconv.FormatInt(id, 10)
}

var responseChanMap sync.Map

func onNewResponse(conn net.Conn, data string) {
	response := &goapplib.Response{}
	err := response.Unmarshal(data)
	if err != nil {
		panic(err)
	}
	ch, ok := responseChanMap.Load(response.TraceId)
	if !ok {
		panic(fmt.Sprintf("responseChanMap.Load(response.TraceId) failed, response.TraceId = %s", response.TraceId))
	}
	ch.(chan *goapplib.Response) <- response
}

func onNewRequest(conn net.Conn, data string) {
	request := &goapplib.Request{}
	err := request.Unmarshal(data)
	if err != nil {
		panic(err)
	}
	switch request.Method {
	case "testCallApp":
		response := &goapplib.Response{
			TraceId: request.TraceId,
			Code:    goapplib.CodeOK,
			Data:    request.Data,
		}
		sendResponse(conn, response)
	}
}

func sendResponse(conn net.Conn, response *goapplib.Response) {
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
	_, err = conn.Write(pack)
	if err != nil {
		log.Printf("[ERRO] client.Write: %v", err)
		return
	}
}

func callLib(conn net.Conn, req *goapplib.Request) *goapplib.Response {
	data := req.Marshal()
	pack, err := dp.DP.Pack(&dp.Message{
		Len:        uint32(len(data)),
		IsResponse: false,
		Data:       data,
	})
	if err != nil {
		panic(err)
	}
	ch := make(chan *goapplib.Response)
	responseChanMap.Store(req.TraceId, ch)
	defer responseChanMap.Delete(req.TraceId)
	_, err = conn.Write(pack)
	if err != nil {
		panic(err)
	}
	response := <-ch
	return response
}

func loopRead(conn net.Conn) {
	for {
		msg, err := dp.DP.Unpack(conn)
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
			onNewResponse(conn, msg.GetData())
		} else {
			//testCallApp
			onNewRequest(conn, msg.GetData())
		}
	}
}

func mockLibCallApp(conn net.Conn) {
	for {
		id := NextTraceId()
		now := time.Now()
		response := goapplib.CallApp(&goapplib.Request{
			TraceId: id,
			Method:  "testCallApp",
			Data:    id,
		})
		if response == nil {
			panic("response is nil")
		}
		if response.Code != goapplib.CodeOK {
			panic(fmt.Sprintf("response.Code != goapplib.CodeOK, response.Code = %d", response.Code))
		}
		if response.Data != id {
			panic(fmt.Sprintf("response.Data != id, response.Data = %s", response.Data))
		}
		log.Printf("[INFO] 测试golib主动调用客户端成功, %s, 耗时: %s", response.TraceId, time.Since(now))
		time.Sleep(time.Second * 5)
	}
}

func mockAppCallLib(conn net.Conn) {
	for {
		id := NextTraceId()
		now := time.Now()
		response := callLib(conn, &goapplib.Request{
			TraceId: id,
			Method:  "test",
			Data:    id,
		})
		if response == nil {
			panic("response is nil")
		}
		if response.Code != goapplib.CodeOK {
			panic(fmt.Sprintf("response.Code != goapplib.CodeOK, response.Code = %d", response.Code))
		}
		if response.Data != id {
			panic(fmt.Sprintf("response.Data != id, response.Data = %s", response.Data))
		}
		log.Printf("[INFO] 测试客户端主动调用golib成功, %s, 耗时: %s", response.TraceId, time.Since(now))
		time.Sleep(time.Second * 6)
	}
}

func main() {
	goapplib.Init(&goapplib.LocalServerConfig{
		Address:  "",
		Callback: new(callback),
	})
	log.Printf("[INFO] goapplib.Address() = %s", goapplib.Address())
	time.Sleep(time.Millisecond * 100)
	conn, err := net.Dial("tcp", goapplib.Address())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	go loopRead(conn)
	go mockLibCallApp(conn) // 模拟golib主动调用客户端
	go mockAppCallLib(conn) // 模拟客户端主动调用golib
	time.Sleep(time.Minute)
}
