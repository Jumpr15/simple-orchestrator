package main

import (
	EntrypointServer "simple-orchestrator/entrypoint"
)

func main() {
	// should init caddy

	go EntrypointServer.StartEntrypointServer()
	
	// should create cluster join endpoint (create a password asw to validate joining machines) => in seperate go routine
	select {}
}