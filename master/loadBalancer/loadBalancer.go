package loadBalancer

import (
	addrType "simple-orchestrator/common/types/addressTypes"

	"net/http"
	"io"
	"log"
	"fmt"
	"bytes"
	"encoding/json"	
)

type CaddyClient struct {
	Address string
	Port string // or int?
}

func (c *CaddyClient) GetUpstreamDetails(node addrType.NodeAddrConfig) {
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
func (c *CaddyClient) Add(node addrType.NodeAddrConfig) { 
	client := &http.Client{}

	req_body := map[string]string{
		"dial": fmt.Sprintf("%s:%s", node.Address, node.Port),
		"@id": node.Id,
	}
	json_req, err := json.Marshal(req_body)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewReqnodest("POST", fmt.Sprintf("http://%s:%s/id/upstream-handler/upstreams/", c.Address, c.Port), bytes.NewBuffer(json_req))
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

func (c *CaddyClient) Delete(node addrType.NodeAddrConfig) {
	client := &http.Client{}

	req, err := http.NewReqnodest("DELETE", fmt.Sprintf("http://%s:%s/id/%s/", c.Address, c.Port, node.Id), nil)
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
