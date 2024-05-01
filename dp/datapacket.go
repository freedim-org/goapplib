package dp

import (
	"bytes"
	"encoding/binary"
	"io"
	"net"
)

type dataPack struct{}

var DP = &dataPack{}

type Message struct {
	Len        uint32
	IsResponse bool
	Data       []byte
}

func (m *Message) GetDataLen() uint32 {
	return m.Len
}

func (m *Message) IsResp() bool {
	return m.IsResponse
}

func (m *Message) GetData() []byte {
	return m.Data
}

// Pack 封包方法(压缩数据)
func (dp *dataPack) Pack(msg *Message) ([]byte, error) {
	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//写dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	//写isResponse
	isResp := int8(0)
	if msg.IsResp() {
		isResp = 1
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, isResp); err != nil {
		return nil, err
	}

	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, []byte(msg.GetData())); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

// Unpack 拆包方法(解压数据)
func (dp *dataPack) Unpack(conn net.Conn) (*Message, error) {
	//先读出dataLen
	headData := make([]byte, 5)
	_, err := io.ReadFull(conn, headData)
	if err != nil {
		return nil, err
	}
	//只解压head的信息，得到dataLen和msgId
	msg := &Message{
		Len: binary.LittleEndian.Uint32(headData[:4]),
	}
	//读isResponse
	msg.IsResponse = headData[4] == 1
	dataTmp := make([]byte, msg.Len)
	//读data数据
	if err := binary.Read(conn, binary.LittleEndian, &dataTmp); err != nil {
		return nil, err
	}
	msg.Data = dataTmp
	return msg, nil
}
