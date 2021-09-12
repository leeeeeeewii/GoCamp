1. 总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用
解包：
fix length
发送端将每个数据包封装为固定长度，不超过缓冲区，接收端每次按固定长度区接受数据

delimiter based
发送端在数据包添加特殊的分隔符，用来标记数据包边界，接收端通过这个边界就可以将不同的数据包拆分开

length field based
发送端在每个数据包头添加包长度信息，接收端在接收到数据后，通过读取包首部的长度字段，获取每一个数据包的实际长度

2. 实现一个从 socket connection 中解码出 goim 协议的解码器。
| parameter     | is required  | type | comment|
| :-----     | :---  | :--- | :---       |
| package length        | true  | int32 bigendian | package length |
| header Length         | true  | int16 bigendian    | header length |
| ver        | true  | int16 bigendian    | Protocol version |
| operation          | true | int32 bigendian | Operation |
| seq         | true | int32 bigendian | jsonp callback |
| body         | false | binary | $(package lenth) - $(header length) |
``` go
package main

import (
	"encoding/binary"
	"errors"
)
type Message struct {
	PLen uint32
	HLen uint16
	Ver  uint16
	Oper uint32
	Seq  uint32
	Body string
}

func Decoder(data []byte) (*Message, error) {
	if len(data) <= 16 {
		return nil, errors.New("data format error")
	}

	m := &Message{}
	m.PLen = binary.BigEndian.Uint32(data[:4])

	m.HLen = binary.BigEndian.Uint16(data[4:6])

	m.Ver = binary.BigEndian.Uint16(data[6:8])

	m.Oper = binary.BigEndian.Uint32(data[8:12])

	m.Seq = binary.BigEndian.Uint32(data[12:16])

	m.Body = string(data[16:])
	return m, nil
}
```
