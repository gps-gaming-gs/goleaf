package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
    {{.imports}}
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		game.GameModule,
		gate.GateModule,
		login.LoginModule,
	)
}