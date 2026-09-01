package healthService

import (
	hbType "simple-orchestrator/common/types/heartbeatTypes"
	addrType "simple-orchestrator/common/types/addressTypes"
	ctrType "simple-orchestrator/common/types/containerTypes"	
	"simple-orchestrator/master/kvStore"
	"simple-orchestrator/master/scheduler"

	"net/rpc"
	"fmt"
	"log"
	"time"
)

func SendHealthcheck(node addrType.NodeAddrConfig, containerConfig ctrType.ContainerConfig) {
	fmt_addr := fmt.Sprintf("[%s]:16767", node.Address) // only works for ipv6
	client, err := rpc.Dial("tcp", fmt_addr) 
	if err != nil {
		log.Println("Error sending healthcheck to worker\n")
		return
	}
	defer client.Close()

	args := hbType.Args{
		NodeAddrConfig: node,
		ContainerConfig: containerConfig,
	}
	var response hbType.Response 

	err = client.Call("HeartbeatServer.SendHealthcheck", args, &response)
	if err != nil {
		log.Println("Error calling heartbeat server rpc method\n")
	}

	fmt.Println("result is:", response, "\n")
	return
}

func StartHealthService(kv *kvStore.KvStore, sch *scheduler.Scheduler)  {
	kvPairs := kv.GetAll() 
	for {
		fmt.Printf("%#v\n\n", kvPairs)
		for _, node := range kvPairs {
			// fmt.Println(key, value, "\n\n")

			containerConfig := sch.ScheduleNode()

			SendHealthcheck(node, containerConfig)
			// send rpc heartbeat requests to endpoints
			// handle 
			// if successful => add endpoint to lb cluster endpoint list
		}


		time.Sleep(3 * time.Second)
	}
}

// should: LOOP: query kv -> heartbeats -> handle accordingly -> update kv + lb