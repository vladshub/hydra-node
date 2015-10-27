package main
import (
	"github.com/hashicorp/mdns"
	"fmt"
)

func main(){
	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			fmt.Printf("Got new entry: %v\n", entry)
		}
	}()

	// Start the lookup
	mdns.Lookup("_hydra._tcp", entriesCh)
	close(entriesCh)
}
