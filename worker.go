package main

import (

)

func main() {
	// should first validate deps (docker)
	// should create healthcheck/state management endpoint in seperate go routine => should split into different goroutines that interact with docker engine to manage container lifecycle
}

func connect() {
	// should connect to cluster master join endpoint with valid passcode
}