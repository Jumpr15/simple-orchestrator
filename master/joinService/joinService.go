package joinService

import (
	"simple-orchestrator/master/kvStore"
	"net"
	"net/http"
	"log"
	"fmt"
)

type NodeAddrConfig struct {
	Address string
	Port string
}

// http handler for onboarding new nodes into cluster
func EntrypointHandler(w http.ResponseWriter, r *http.Request) {
	address, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Fatal(err)
	}

	node := NodeAddrConfig{
		Address: address,
		Port: port,
	}

	fmt.Fprintf(w, "%+v", node)
	// append node to kv store
}

func StartJoinService(kv *kvStore.KVStore) {
	log.Printf("Starting join service...")

	http.HandleFunc("/join", EntrypointHandler)

	http.ListenAndServe(":8090", nil)

}