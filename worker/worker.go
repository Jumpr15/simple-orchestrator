package main

import (
	"simple-orchestrator/worker/listenerService"
	"simple-orchestrator/worker/containerDaemonService"

	"github.com/dgraph-io/ristretto/v2"

	"net/http"
	"log"
)
func main() {
	cfg := ristretto.Config[string, any]{ // [key, value] (types)
		NumCounters: 1e7,
		MaxCost: 1 << 30,
		BufferItems: 64,
	}
	nodeKv, err := ristretto.NewCache(&cfg)
	if err != nil {
		log.Fatal("Failed creating cache")
	}
	defer nodeKv.Close()

	connect()
	// 
	containerDaemonService.StartDaemonService()
	listenerService.StartListenerServer(nodeKv)
}

func connect() {
	master_addr := "http://localhost:8090/join"
	client := http.Client{}
	res, err := client.Get(master_addr)
	if err != nil {
		log.Fatal("Error connecting to master")
	}
	log.Println("Sent conn request to master", res)
}