package loadBalancer

import (
	addrType "simple-orchestrator/common/types/addressTypes"
	lbType "simple-orchestrator/common/types/lbConfigTypes"

	"net/http"
	"io"
	"log"
	"fmt"
	"bytes"
	"encoding/json"	
)

func New(Address string, Port string) *lbType.LBConfigClient {
	lb := lbType.LBConfigClient{
		Address: Address,
		Port: Port,
	}
	return &lb
}

func (lb *lbType.LBConfigClient) GetUpstreamDetails(node addrType.NodeAddrConfig) {
	res, err := http.Get(fmt.Sprintf("http://%s:%s/id/%s/", c.Address, c.Port, node.Id))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	fmt.Printf("%s\n", body)
}

// assumes http routes handler has @id: "upstream-handler"]
// sends json payload of { "dial": address, "@id": id }
func (lb *lbType.LBConfigClient) Add(node addrType.NodeAddrConfig) { 
	client := &http.Client{}

	req_body := map[string]string{
		"dial": fmt.Sprintf("%s:%s", node.Address, node.Port),
		"@id": node.Id,
	}
	json_req, err := json.Marshal(req_body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s:%s/id/upstream-handler/upstreams/", c.Address, c.Port), bytes.NewBuffer(json_req))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	fmt.Printf("%s\n", body)
}

func (lb *lbType.LBConfigClient) Delete(node addrType.NodeAddrConfig) {
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s:%s/id/%s/", c.Address, c.Port, node.Id), nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	fmt.Printf("%s\n", body)
}
