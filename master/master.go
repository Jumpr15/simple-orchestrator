package main

import (
	"simple-orchestrator/master/joinService"  
	"simple-orchestrator/master/healthService"

	"simple-orchestrator/master/scheduler"
	"simple-orchestrator/master/kvStore"

	"context"
	"os/signal"
	"syscall"

	"fmt"
)

func main() {
	// should init kv store + lb

	nodeKv := kvStore.New() // init with config details // change to cluster kv?
	sch := scheduler.New(nodeKv)
	loadBalancer := 

	go healthService.StartHealthService(nodeKv, sch) // takes in kv store, lb / eventually consistent health checks


	// should create cluster join endpoint (create a password asw to validate joining machines) => in seperate go routine
	go joinService.StartJoinService(nodeKv) // takes in kv store


	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM,
	)
	defer stop()
	<-ctx.Done()
	fmt.Println("Exiting")
}