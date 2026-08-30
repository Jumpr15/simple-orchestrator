package main

import (
	"simple-orchestrator/worker/listenerService"

	"net/http"
	"log"
)
func main() {
	// should first validate deps (docker)
	// should create healthcheck/state management endpoint in seperate go routine => should split into different goroutines that interact with docker engine to manage container lifecycle
	connect()
	listenerService.StartListenerServer()

	// if connect is successful
	// handle connections
}

func connect() {
	// fixed value => in real app check port availability and send port num? maybe server auto extracts
	master_addr := "http://localhost:8090/join"
	client := http.Client{}
	res, err := client.Get(master_addr)
	if err != nil {
		log.Fatal("Error connecting to master")
	}
	log.Println("Sent conn request to master", res)

	// should take master: addr, port, fixed endpoint
	// handle for errs
	//// should connect to cluster master join endpoint with valid passcode
}