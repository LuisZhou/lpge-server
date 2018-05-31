package main

import (
	_ "encoding/binary"
	"fmt"
	"github.com/LuisZhou/lpge/network"
	"github.com/LuisZhou/lpge/network/processor/protobuf"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}

	// Hello message (JSON-encoded)
	// The structure of the message
	// data := []byte(`{
	//        "Hello": {
	//            "Name": "leaf"
	//        }
	//    }`)

	msgParser := network.NewMsgParser()
	msgParser.SetByteOrder(true)

	person := &protobuf.Person{}
	person.Name = "abc"

	p := protobuf.NewProtobufProcessor()

	p.Register(1, protobuf.Person{})
	data, err1 := p.Marshal(1, person)
	if err1 != nil {
		panic(err1)
	}

	// len + data
	// m := make([]byte, 4+len(data))

	// binary.BigEndian.PutUint16(m, uint16(1))

	// // BigEndian encoded
	// binary.BigEndian.PutUint16(m+2, uint16(len(data)))

	// copy(m[4:], data)

	// // Send message
	// conn.Write(m)

	msgParser.Write(conn, 1, data)

	cmd, ret, err2 := msgParser.Read(conn)

	fmt.Println(cmd, ret, err2)

	time.Sleep(time.Second * 5)

}
