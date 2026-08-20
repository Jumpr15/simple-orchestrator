package caddy

import (
	"net/http"
	"io"
	"log"
	"fmt"
	"bytes"
	"encoding/json"
)

// this should be a interface? for flexibility between different caddy modules (normal vs admin etc)
type CaddyClient struct {
	Address string
	Port string // or int?
}

type UpstreamEndpoint struct {
	Id string
	Address string
	Port string
}

func (c *CaddyClient) GetUpstreamDetails(ue UpstreamEndpoint) {
	res, err := http.Get(fmt.Sprintf("http://%s:%s/id/%s/", c.Address, c.Port, ue.Id))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	fmt.Printf("%s\n", body)
}

// assumes http routes handler has @id: "upstream-handler"]
// sends json payload of { "dial": address, "@id": id }
func (c *CaddyClient) AddUpstreamEndpoint(ue UpstreamEndpoint) { 
	client := &http.Client{}

	req_body := map[string]string{
		"dial": fmt.Sprintf("%s:%s", ue.Address, ue.Port),
		"@id": ue.Id,
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

func (c *CaddyClient) DeleteUpstreamEndpoint(ue UpstreamEndpoint) {
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s:%s/id/%s/", c.Address, c.Port, ue.Id), nil)
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