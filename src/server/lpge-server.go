package main

import (
	"github.com/LuisZhou/lpge/conf"
	"github.com/LuisZhou/lpge/gate"
	"github.com/LuisZhou/lpge/log"
	"github.com/LuisZhou/lpge/module"
	"os"
	"os/signal"
	sconf "server/conf"
	server_gate "server/gate"
)

//func Run(mods ...module.Module) {
func Run(mods map[string]module.Module) {
	// logger
	if conf.LogLevel != "" {
		logger, err := log.New(conf.LogLevel, conf.LogPath, conf.LogFlag)
		if err != nil {
			panic(err)
		}
		log.Export(logger)
		defer logger.Close()
	}

	log.Release("LPGE %v starting up", 1.0)

	// module
	for k, v := range mods {
		module.Register(v, k)
	}

	// cluster
	//cluster.Init()

	// console
	//console.Init()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Release("LPGE closing down (signal: %v)", sig)
	//console.Destroy()
	//cluster.Destroy()
	module.Destroy()
}

func main() {
	conf.LogLevel = sconf.Server.LogLevel
	conf.LogPath = sconf.Server.LogPath
	conf.LogFlag = sconf.LogFlag
	conf.ConsolePort = sconf.Server.ConsolePort
	conf.ProfilePath = sconf.Server.ProfilePath

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

	Run(map[string]module.Module{
		"gate": _gate,
	})
}
