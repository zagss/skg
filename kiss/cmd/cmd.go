package main

import (
	"kiss/config"
	"kiss/router"
)

func main() {
	config.InitConfig()
	r := router.NewRouter()
	_ = r.Run(config.Config.System.HttpPort)
}
