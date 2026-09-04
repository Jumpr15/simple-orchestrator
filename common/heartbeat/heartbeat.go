package heartbeat

import (
	hbType "simple-orchestrator/common/types/heartbeatTypes"
	ctrType "simple-orchestrator/common/types/containerTypes"
	// addrType "simple-orchestrator/common/types/addressTypes"
	"simple-orchestrator/worker/containerService"

	"github.com/dgraph-io/ristretto/v2"

	"log"
)

type HeartbeatServer struct {
	Kv *ristretto.Cache[string, any] 
}

func (hbs *HeartbeatServer) SendHealthcheck(args *hbType.Args, res *hbType.Response) error {
	var currentNumContainers int
	currentNumContainersInterface, found := hbs.Kv.Get("currentNumContainers")
	if !found {
		log.Println("currentNumContainers not set")
		hbs.Kv.Set("currentNumContainers", 0, 1) 
		hbs.Kv.Wait()
		
		currentNumContainers = 0
	} else {
		var ok bool
		currentNumContainers, ok = currentNumContainersInterface.(int)
		if !ok {
			log.Fatal("currentNumContainers is not a int?")
		}
	}

	desiredNumContainers := args.DesiredNumContainers

	if currentNumContainers > desiredNumContainers {
		log.Printf("Too many containers, destroying containers")
	}
	if currentNumContainers < desiredNumContainers {
		log.Printf("Not enough containers, spinning up new containers")
		// loop for diff between curr and desired
		for i := 0; i < desiredNumContainers-currentNumContainers; i++ { 
			go containerService.CreateAndStartContainer(args.ContainerConfig)
		}
	}
	if currentNumContainers == desiredNumContainers {
		log.Printf("Just enough...for now, doing nothing")
	}

	*res = hbType.Response{
		NodeAddrConfig: args.NodeAddrConfig, // resend with sender args
		ContainerList: ctrType.ContainerList{}, // query from kv
		CurrentNumContainers: currentNumContainers, // query from kv
	}
	return nil
}