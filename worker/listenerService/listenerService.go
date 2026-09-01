package listenerService

import (
	"simple-orchestrator/common/heartbeat"

	"github.com/dgraph-io/ristretto/v2"

	"log"
	"net"
	"net/rpc"
)

func StartListenerServer(kv *ristretto.Cache[string, any]) {
	server := heartbeat.HeartbeatServer{
		KvStore: kv,
	}
	rpc.Register(&server)
	
	listener, err := net.Listen("tcp", ":16767")
	if err != nil {
		log.Fatal("Listener RPC server failed\n")
	}
	defer listener.Close()
	
	log.Println("Listener server starting...\n")
	rpc.Accept(listener)
}

