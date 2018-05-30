package main

import (
	"github.com/LuisZhou/lpge"
	"github.com/LuisZhou/lpge/conf"
	"github.com/LuisZhou/lpge/gate"
	"github.com/LuisZhou/lpge/log"
	"github.com/LuisZhou/lpge/module"
	sconf "server/conf"
	server_gate "server/gate"
)

func main() {
	conf.LogLevel = sconf.Server.LogLevel
	conf.LogPath = sconf.Server.LogPath
	conf.LogFlag = sconf.LogFlag
	conf.ConsolePort = sconf.Server.ConsolePort
	conf.ProfilePath = sconf.Server.ProfilePath

	conf.GateConfig = sconf.GateConfig
	conf.AgentConfig = sconf.AgentConfig
	conf.FunctionConfig = sconf.FunctionConfig

	log.Debug(sconf.Server.TCPAddr)

	_gate := &gate.Gate{
		MaxConnNum:      sconf.Server.MaxConnNum,
		PendingWriteNum: sconf.PendingWriteNum,
		MaxMsgLen:       sconf.MaxMsgLen,
		HTTPTimeout:     sconf.HTTPTimeout,
		TCPAddr:         sconf.Server.TCPAddr,
		WSAddr:          sconf.Server.WSAddr,
		LittleEndian:    sconf.LittleEndian,
		NewWsAgent:      server_gate.NewWsAgent,
		NewTcpAgent:     server_gate.NewTcpAgent,
	}

	lpge.Run(map[string]module.Module{
		"gate": _gate,
	})
}
