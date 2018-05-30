package gate

import (
	"github.com/LuisZhou/lpge/gate"
	"github.com/LuisZhou/lpge/network"
	"github.com/LuisZhou/lpge/network/processor/protobuf"
)

var (
	Processor network.Processor
	handlers  map[interface{}]interface{}
)

func init() {
	handlers = make(map[interface{}]interface{})
	// the first parameter must the agent.
	handlers[uint16(1)] = func(args []interface{}) (ret interface{}, err error) {
		return nil, nil
	}

	person := &protobuf.Person{}
	person.Name = "abc"

	Processor = protobuf.NewProtobufProcessor()
	Processor.Register(1, protobuf.Person{})
}

type NewAgent struct {
	gate.AgentTemplate
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
