package heartbeat

import (
	hbType "simple-orchestrator/common/types/heartbeatTypes"
	ctrType "simple-orchestrator/common/types/containerTypes"
	// addrType "simple-orchestrator/common/types/addressTypes"

	"github.com/dgraph-io/ristretto/v2"
)

type HeartbeatServer struct {
	KvStore *ristretto.Cache[string, any] 
}

func (hbs *HeartbeatServer) SendHealthcheck(args *hbType.Args, res *hbType.Response) error {
	*res = hbType.Response{
		NodeAddrConfig: args.NodeAddrConfig,
		ContainerList: ctrType.ContainerList {},
		CurrentNumContainers: 2, // query from kv
	}
	return nil
}