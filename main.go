package main

import (
	//	"fmt"
	"github.com/vladshub/hydra-node/players"
	"github.com/vladshub/hydra-node/discovery"
	"github.com/vladshub/hydra-node/control"
)

func main() {

	// 1. Register to mDns, consul, etcd ....
	// 2. Create an control interface
	// 2.1.  gRPC interface
	// 3. File path validations
	// 3.1. Mount ability
	// 5. Create player
	// 5.1. omxplayer
	// 6. Attach player to control interface

	// Setup our service export
	dir, err := discovery.DiscoveryFuctory("nDNS", "Hydra Client", "Hydra")
	if err != nil {
		panic(err)
	}
	err = dir.Register()

	if err != nil {
		panic(err)
	}

	player, err := players.PlayerFactory("OmxPlayer")
	if err != nil {
		panic(err)
	}

	control, err := control.ControlFactory("HttpControl", player, nil)
	if err != nil {
		panic(err)
	}

	control.Listen()


	// Setup HTTP server
	//	router := gin.Default()

}
