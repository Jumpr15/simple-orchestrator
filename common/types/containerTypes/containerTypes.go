package containerTypes 

import (
	"net/netip"

	// "github.com/moby/moby/api/types/network/v2"
)

type PortConfig struct {
 	AddressString string 
 	HostInt uint16 // represented by int type
 	ExposedInt uint16 // represented by int type
 	// Protocol network.IPProtocol // represented by string type
 	Address netip.Addr
	// HostPort network.Port
	// ExposedPort network.Port
}

type ContainerConfig struct {
	DesiredNumContainers int
	ImageName string
	EnvMap map[string]int // change to string any later
	PortConfig
}

type ContainerState struct {
	ID string
	State string // enum
}

type ContainerList []ContainerState