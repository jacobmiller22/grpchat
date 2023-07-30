package main

import (
	"flag"

	"github.com/jacobmiller22/grpchat/client"
	"github.com/jacobmiller22/grpchat/server"
)

var serverFlag = flag.Bool("is_server", false, "Whether or not to run this application as the server application")

func main() {
	flag.Parse()

	if *serverFlag == true {
		server.Main()
	} else {
		client.Main()
	}
}
