package main

import (
	"fmt"
	"github.com/LuisZhou/lpge/conf"
	"github.com/LuisZhou/lpge/gate"
	"github.com/LuisZhou/lpge/log"
	"github.com/LuisZhou/lpge/module"
	"os"
	"os/signal"
	server_gate "server/gate"
	"time"
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
	//module.Init()

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
	fmt.Println(conf.LenStackBuf)
	_gate := &gate.Gate{
		MaxConnNum:      2,
		PendingWriteNum: 20,
		MaxMsgLen:       4096, // todo: double check.
		HTTPTimeout:     10 * time.Second,
		TCPAddr:         "0.0.0.0:3563",
		WSAddr:          "0.0.0.0:6001",
		LittleEndian:    true,
		NewWsAgent:      server_gate.NewWsAgent,
		NewTcpAgent:     server_gate.NewTcpAgent,
	}

	Run(map[string]module.Module{
		"gate": _gate,
	})
}
