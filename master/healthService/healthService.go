package healthService

import (
	"simple-orchestrator/common/heartbeat"
	"simple-orchestrator/master/kvStore"

	"net/rpc"
	"fmt"
	"log"
	"time"
)

func SendHealthcheck() {
	client, err := rpc.Dial("tcp", "localhost:16767")
	if err != nil {
		log.Println("Error sending healthcheck to worker\n")
		return
	}
	defer client.Close()

	args := heartbeat.Args{}
	var response heartbeat.Response 

	err = client.Call("HeartbeatServer.SendHealthcheck", args, &response)
	if err != nil {
		log.Println("Error calling heartbeat server rpc method\n")
	}

	fmt.Println("result is:", response, "\n")
	return
}

func StartHealthService(kv *kvStore.KvStore)  {
	kvPairs := kv.GetAll() // type map[stirng]any
	for {
		fmt.Printf("%#v\n\n", kvPairs)
		for key, value := range kvPairs {
			fmt.Println(key, value, "\n\n")
			SendHealthcheck()
			// send rpc heartbeat requests to endpoints
			// handle 
			// if successful => add endpoint to lb cluster endpoint list
		}


		time.Sleep(3 * time.Second)
	}
}

// should: LOOP: query kv -> heartbeats -> handle accordingly -> update kv + lb