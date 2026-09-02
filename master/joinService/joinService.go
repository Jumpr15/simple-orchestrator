package joinService

import (
	addrType "simple-orchestrator/common/types/addressTypes"
	"simple-orchestrator/master/kvStore"
	
	"net"
	// "net/netip"
	"net/http"
	"uuid"
	"log"
	// "fmt"
)

type EntrypointServer struct {
	kv *kvStore.KvStore
}

// http handler for onboarding new nodes into cluster
func (es *EntrypointServer) EntrypointHandler(w http.ResponseWriter, r *http.Request) {
	addr, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Fatal(err)
	}

	// addr, err := netip.ParseAddr(host)

	// full_addr := addr.StringExpanded()

	id := uuid.New()
	id_string := id.String()

	node := addrType.NodeAddrConfig{
		Address: addr,
		Port: port,
		Id: id_string,
	}

	// call scheduler interface and allocate based on alloc-algorithm 
	

	es.kv.Set(id_string, node)


	// val, found := es.kv.Get(id_string)
	// fmt.Fprintf(w, "%+v", node)
	// fmt.Println(val, found, id_string)
}

func StartJoinService(kv *kvStore.KvStore) {
	log.Printf("Starting join service...\n\n")

	es := EntrypointServer{kv}

	http.HandleFunc("/join", es.EntrypointHandler)

	http.ListenAndServe(":8090", nil)

}