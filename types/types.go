package types

import (
	cnitypes "github.com/containernetworking/cni/pkg/types"
	"net"
)

// Net is The top-level network config - IPAM plugins are passed the full configuration
// of the calling plugin, not just the IPAM section.
type Net struct {
	Name       string      `json:"name"`
	CNIVersion string      `json:"cniVersion"`
	IPAM       *IPAMConfig `json:"ipam"`
}

// IPAMConfig describes the expected json configuration for this plugin
type IPAMConfig struct {
	Name       string
	Type       string `json:"type"`
	Range      string `json:"range"`
	GatewayStr string `json:"gateway"`
	Gateway    net.IP
	EtcdHost   string            `json:"etcd_host"`
	Routes     []*cnitypes.Route `json:"routes"`
	DNS        cnitypes.DNS      `json:"dns"`
}

// IPAMEnvArgs are the environment vars we expect
type IPAMEnvArgs struct {
	cnitypes.CommonArgs
	IP      cnitypes.UnmarshallableString `json:"ip,omitempty"`
	GATEWAY cnitypes.UnmarshallableString `json:"gateway,omitempty"`
}
