package discovery
import "errors"

type Discovery interface {
	Register() error
}

func DiscoveryFuctory(discovery_type string, component_name string, service_name string) (Discovery, error) {
	switch discovery_type {
	case "mDNS":
		return NewMdnsServer(component_name, service_name)
	}
	return nil, errors.New("Unknown discovery service")
}
