package main

import (
	"fmt"
	"github.com/vladshub/hydra-node/control"
	"github.com/vladshub/hydra-node/discovery"
	"github.com/vladshub/hydra-node/players"
)

func main() {

	// 1. Register to mDns, consul, etcd ....
	// 2. Create an control interface
	// 2.1.  gRPC interface
	// 2.2.  HTTP interface
	// 3. File System Manager
	// 3.1. Mount ability
	// 3.1.1. NFS
	// 3.1.1. SMB
	// 3.1.1. ?
	// 5. Create player
	// 5.1. omxplayer
	// 6. Attach player to control interface

	// Setup our service export
	disc, err := discovery.DiscoveryFuctory("mDNS", "Hydra Client", "Hydra")
	if err != nil {
		panic(err)
	}
	err = disc.Register()

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

	control.Listen("0.0.0.0:8080")
	fmt.Println("Done")
	// Setup HTTP server
	//	router := gin.Default()

}
