package healthService

import (
	"simple-orchestrator/master/kvStore"

	"fmt"
	"time"
)

func StartHealthService(kv *kvStore.KVStore)  {
	kvPairs := kv.GetAll() // type map[stirng]any
	for {
		fmt.Printf("%#v\n", kvPairs)
		for key, value := range kvPairs {
			fmt.Println(key, value)
			// send rpc heartbeat requests to endpoints
			// handle 
			// if successful => add endpoint to lb cluster endpoint list
		}


		time.Sleep(3 * time.Second)
	}
}

// should: LOOP: query kv -> heartbeats -> handle accordingly -> update kv + lb