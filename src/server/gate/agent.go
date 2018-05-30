package gate

import (
	"fmt"
	"github.com/LuisZhou/lpge/gate"
	_ "github.com/LuisZhou/lpge/log"
	"github.com/LuisZhou/lpge/network"
	"github.com/LuisZhou/lpge/network/processor/protobuf"
)

type NewAgent struct {
	gate.AgentTemplate
}

var (
	Processor network.Processor
	handlers  map[interface{}]interface{}
)

func init() {
	handlers = make(map[interface{}]interface{})
	// the first parameter must the agent.
	handlers[uint16(1)] = func(args []interface{}) (ret interface{}, err error) {
		fmt.Println(args[0])
		a := args[0].(*gate.AgentTemplate)
		a.WriteMsg(uint16(2), args[1])
		return nil, nil
	}

	person := &protobuf.Person{}
	person.Name = "abc"

	Processor = protobuf.NewProtobufProcessor()
	Processor.Register(1, protobuf.Person{})
}

func NewWsAgent(conn network.Conn, gate *gate.Gate) network.Agent {
	a := &NewAgent{}
	a.Init(conn, gate)
	a.Skeleton.SetChanRPCHandlers(handlers)
	a.Processor = Processor
	return a
}

func NewTcpAgent(conn network.Conn, gate *gate.Gate) network.Agent {
	a := &NewAgent{}
	a.Init(conn, gate)
	a.Skeleton.SetChanRPCHandlers(handlers)
	a.Processor = Processor
	return a
}
