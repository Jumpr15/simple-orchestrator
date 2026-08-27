package main

import (
	"simple-orchestrator/master/joinService"  
	"simple-orchestrator/master/kvStore"
	"simple-orchestrator/master/healthService"
)

func main() {
	// should init kv store + lb

	// kv := kvStore.New(...) // init with config details
	// lb := ... // (implement caddy init)

	// should create cluster join endpoint (create a password asw to validate joining machines) => in seperate go routine
	go joinService.StartJoinService() // takes in kv store

	go healthservice.StartHealthService() // takes in kv store, lb
	
	select {} // check for sigint
}