package entrypoint

import (
	"net"
	"net/http"
	"log"
	"fmt"
)

type NodeConfig struct {
	Address string
	Port string
}

func EntrypointHandler(w http.ResponseWriter, r *http.Request) {
	address, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Fatal(err)
	}

	node := NodeConfig{
		Address: address,
		Port: port,
	}

	fmt.Fprintf(w, "%+v", node)
}

func StartEntrypointServer() {
	log.Printf("Starting entrypoint server...")

	http.HandleFunc("/entrypoint", EntrypointHandler)

	http.ListenAndServe(":8090", nil)
}