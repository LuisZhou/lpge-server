package conf

import (
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

	// todo: seperate to three catalog:
	// 1. gate
	// 2. agent
	// 3. function module.

	// skeleton conf
	GoLen              = 10000
	TimerDispatcherLen = 10000
	AsynCallLen        = 10000
	ChanRPCLen         = 10000
)
