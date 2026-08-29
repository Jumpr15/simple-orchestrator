package main

import (
	"simple-orchestrator/master/joinService"  
	"simple-orchestrator/master/kvStore"
	"simple-orchestrator/master/healthService"

	"context"
	"os/signal"
	"syscall"

	"fmt"
)

func main() {
	// should init kv store + lb

	nodeKv := kvStore.New() // init with config details
	// lb := ... // (implement caddy init)

	// should create cluster join endpoint (create a password asw to validate joining machines) => in seperate go routine
	go joinService.StartJoinService(nodeKv) // takes in kv store

	go healthService.StartHealthService(nodeKv) // takes in kv store, lb / eventually consistent health checks

	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM,
	)
	defer stop()
	<-ctx.Done()
	fmt.Println("Exiting")
}