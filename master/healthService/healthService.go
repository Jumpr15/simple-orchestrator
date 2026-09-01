package healthService

import (
	hbType "simple-orchestrator/common/types/heartbeatTypes"
	"simple-orchestrator/master/kvStore"

	"net/rpc"
	"fmt"
	"log"
	"time"
)

func SendHealthcheck(ipv6_addr string) {
	fmt_addr := fmt.Sprintf("[%s]:16767", ipv6_addr)
	client, err := rpc.Dial("tcp", fmt_addr) // for handling ipv6 addrs
	if err != nil {
		log.Println("Error sending healthcheck to worker\n")
		return
	}
	defer client.Close()

	args := hbType.Args{}
	var response hbType.Response 

	err = client.Call("HeartbeatServer.SendHealthcheck", args, &response)
	if err != nil {
		log.Println("Error calling heartbeat server rpc method\n")
	}

	fmt.Println("result is:", response, "\n")
	return
}

func StartHealthService(kv *kvStore.KvStore)  {
	kvPairs := kv.GetAll() 
	for {
		fmt.Printf("%#v\n\n", kvPairs)
		for _, node := range kvPairs {
			// fmt.Println(key, value, "\n\n")
			SendHealthcheck(node.Address)
			// send rpc heartbeat requests to endpoints
			// handle 
			// if successful => add endpoint to lb cluster endpoint list
		}


		time.Sleep(3 * time.Second)
	}
}

// should: LOOP: query kv -> heartbeats -> handle accordingly -> update kv + lb