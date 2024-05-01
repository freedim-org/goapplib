package goapplib

import (
	"google.golang.org/protobuf/proto"
)

func (r *GoRequest) Unmarshal(data []byte) error {
	return proto.Unmarshal([]byte(data), r)
}

func (r *GoRequest) Marshal() []byte {
	data, _ := proto.Marshal(r)
	return data
}

func (r *GoResponse) Marshal() []byte {
	data, _ := proto.Marshal(r)
	return data
}

func (r *GoResponse) Unmarshal(data []byte) error {
	return proto.Unmarshal([]byte(data), r)
}
