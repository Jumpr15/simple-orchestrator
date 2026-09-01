package heartbeat

import (
	hbType "simple-orchestrator/common/types/heartbeatTypes"
	ctrType "simple-orchestrator/common/types/containerTypes"
	addrType "simple-orchestrator/common/types/addressTypes"
)

type HeartbeatServer struct {

}

func (hbs *HeartbeatServer) SendHealthcheck(args *hbType.Args, res *hbType.Response) error {
	*res = hbType.Response{
		NodeAddrConfig: args.NodeAddrConfig,
		ContainerList: ctrType.ContainerList {},
		CurrentNumContainers: 2,
	}
	return nil
}