package main

import (
	"flag"
	"github.com/Leorevoir/Go-IRCD/pkg/server"
	"runtime"
)

var config_file = flag.String("config", "./config/ircd.json", "Path to the configuration file for the IRC daemon's configuration")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	config := server.FromJSON(*config_file)
	server.Run(config)
}
