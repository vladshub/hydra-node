package discovery

import (
	"github.com/hashicorp/mdns"
	"os"
)

type Mdns struct {
	componentName string
	serviceName   string
}

func NewMdnsServer(component_name string, service_name string) (*Mdns, error) {
	return Mdns{
		componentName: component_name,
		serviceName:   service_name,
	}
}

func (dir *Mdns) Register() error {
	host, _ := os.Hostname()
	service, _ := mdns.NewMDNSService(host,
		"_client._hydra._tcp",
		"",
		host,
		8000,
		nil,
		[]string{dir.serviceName, dir.componentName},
	)

	// Create the mDNS server, defer shutdown
	server, err := mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		return err
	}
	defer server.Shutdown()
	return nil
}
