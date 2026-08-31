package listenerService

import (
	"simple-orchestrator/common/heartbeat"

	"log"
	"net"
	"net/rpc"
)

func StartListenerServer() {
	server := new(heartbeat.HeartbeatServer)
	rpc.Register(server)
	
	listener, err := net.Listen("tcp", ":16767")
	if err != nil {
		log.Fatal("Listener RPC server failed\n")
	}
	defer listener.Close()
	
	log.Println("Listener server starting...\n")
	rpc.Accept(listener)
}

