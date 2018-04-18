package conf

import (
	"github.com/LuisZhou/lpge/conf"
	"log"
	"time"
)

var (
	// log conf
	LogFlag = log.LstdFlags

	// gate conf
	PendingWriteNum        = 2000
	MaxMsgLen       uint16 = 4096 // double check
	HTTPTimeout            = 10 * time.Second
	LenMsgLen              = 2
	LittleEndian           = true

	// gate conf
	GateConfig = conf.ModuleConfig{
		GoLen:              10,
		TimerDispatcherLen: 10,
		AsynCallLen:        10,
		ChanRPCLen:         10,
		TimeoutAsynRet:     10,
	}

	// agent config
	AgentConfig = conf.ModuleConfig{
		GoLen:              10,
		TimerDispatcherLen: 10,
		AsynCallLen:        100,
		ChanRPCLen:         100,
		TimeoutAsynRet:     10,
	}

	// function module config
	FunctionConfig = conf.ModuleConfig{
		GoLen:              10,
		TimerDispatcherLen: 10,
		AsynCallLen:        1000,
		ChanRPCLen:         1000,
		TimeoutAsynRet:     10,
	}
)
