package joinService

import (
	"simple-orchestrator/master/kvStore"
	"net"
	"net/http"
	"uuid"
	"log"
	"fmt"
)

type EntrypointServer struct {
	kv *kvStore.KVStore
}

type NodeAddrConfig struct {
	Address string
	Port string
}

// http handler for onboarding new nodes into cluster
func (es *EntrypointServer) EntrypointHandler(w http.ResponseWriter, r *http.Request) {
	address, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Fatal(err)
	}

	node := NodeAddrConfig{
		Address: address,
		Port: port,
	}

	id := uuid.New()
	id_string := id.String()
	es.kv.Set(id_string, node)


	val, found := es.kv.Get(id_string)
	fmt.Fprintf(w, "%+v", node)
	fmt.Println(val, found, id_string)
}

func StartJoinService(kv *kvStore.KVStore) {
	log.Printf("Starting join service...")

	es := EntrypointServer{kv}

	http.HandleFunc("/join", es.EntrypointHandler)

	http.ListenAndServe(":8090", nil)

}