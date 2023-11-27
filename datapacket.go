package goapplib

import (
	"bytes"
	"encoding/binary"
	"io"
	"net"
)

type dataPack struct{}

var DP = &dataPack{}

type DPMessage struct {
	len  uint32
	data string
}

func (m *DPMessage) GetDataLen() uint32 {
	return m.len
}

func (m *DPMessage) GetData() string {
	return m.data
}

// Pack 封包方法(压缩数据)
func (dp *dataPack) Pack(msg *DPMessage) ([]byte, error) {
	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//写dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, []byte(msg.GetData())); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

// Unpack 拆包方法(解压数据)
func (dp *dataPack) Unpack(conn net.Conn) (*DPMessage, error) {
	//先读出dataLen
	headData := make([]byte, 4)
	_, err := io.ReadFull(conn, headData)
	if err != nil {
		return nil, err
	}
	//只解压head的信息，得到dataLen和msgId
	msg := &DPMessage{
		len: binary.LittleEndian.Uint32(headData),
	}
	dataTmp := make([]byte, msg.len)
	//读data数据
	if err := binary.Read(conn, binary.LittleEndian, &dataTmp); err != nil {
		return nil, err
	}
	msg.data = string(dataTmp)
	return msg, nil
}
