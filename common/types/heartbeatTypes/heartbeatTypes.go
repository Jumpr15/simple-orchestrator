package heartbeatTypes

import (
	ctrType "simple-orchestrator/common/types/containerTypes"
	addrType "simple-orchestrator/common/types/addressTypes"
)

type Args struct {
	addrType.NodeAddrConfig
	ctrType.ContainerConfig
}

type Response struct {
	addrType.NodeAddrConfig
	ctrType.ContainerList 

	CurrentNumContainers int // Needed?
}
