package discovery

import (
	"github.com/hashicorp/mdns"
	"fmt"
)

type Mdns struct {
	componentName string
	serviceName   string
}

func NewMdnsServer(component_name string, service_name string) (*Mdns, error) {
	return &Mdns{
		componentName: component_name,
		serviceName:   service_name,
	}, nil

}

func (dir *Mdns) Register() error {
	service, err := mdns.NewMDNSService("hydra_client",
		"_hydra._tcp",
		"",
		"",
		8000,
		nil,
		[]string{dir.serviceName, dir.componentName},
	)
	fmt.Println(service, err)
	// Create the mDNS server, defer shutdown
	server, err := mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		return err
	}
	defer server.Shutdown()
	return nil
}
