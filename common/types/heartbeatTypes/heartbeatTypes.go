package heartbeatTypes

import (
	ctrType "simple-orchestrator/common/types/containerTypes"
	addrType "simple-orchestrator/common/types/addressTypes"
)

type Args struct {
	addrType.NodeAddrConfig
	ctrType.ContainerConfig

	DesiredNumContainers int // Needed? (Running or Total )
}

type Response struct {
	addrType.NodeAddrConfig
	ctrType.ContainerList 

	CurrentNumContainers int // Needed?
}
