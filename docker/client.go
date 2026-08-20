package docker

import (
	"context"
	"log"
	"net/netip"
	"fmt"
	"strconv"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/network"
	"github.com/moby/moby/client"
)

// converts map of ENV_VAR_NAME: INT_VALUE to passable array into docker env config as arg
func MapToEnvArr(envMap map[string]int) []string {
	var envArr []string
	var envVar string
	for k, v := range envMap {
		envVar = fmt.Sprintf("%s=%s", k, strconv.Itoa(v))
		envArr = append(envArr, envVar)
	}
	return envArr
}

type PortConfig struct {
	AddressString string 
	HostInt uint16 // represented by int type
	ExposedInt uint16 // represented by int type
	Protocol network.IPProtocol // represented by string type
	Address netip.Addr
	HostPort network.Port
	ExposedPort network.Port
}

// converts string list of ports to expose for a container to docker Config/HostConfig compatible types to pass as args
func ConvertPortConfig(portConfigList []PortConfig) (network.PortSet, network.PortMap) { 
	for i := range portConfigList {
		portConfigAddress, err := netip.ParseAddr(portConfigList[i].AddressString)
		if err != nil {
			log.Fatal(err)
		}
		portConfigList[i].Address = portConfigAddress

		portConfigHostPort, ok := network.PortFrom(portConfigList[i].HostInt, portConfigList[i].Protocol)
		if ok == false  {
			log.Fatal("Error occurred during converting HostInt to HostPort")
		}
		portConfigList[i].HostPort = portConfigHostPort

		portConfigExposedPort, ok := network.PortFrom(portConfigList[i].ExposedInt, portConfigList[i].Protocol)	
		if ok == false {
			log.Fatal("Error occurred during converting ExposedInt to ExposedPort")
		}
		portConfigList[i].ExposedPort = portConfigExposedPort
	}
	fmt.Println(portConfigList)

	exposedPorts := make(map[network.Port]struct{}) // map of type network.PortSet
	portBindings := make(map[network.Port][]network.PortBinding) // map of type network.PortMap

	for _, portConfig := range portConfigList {
		exposedPorts[portConfig.ExposedPort] = struct{}{} 
		portBindings[portConfig.HostPort] = []network.PortBinding{{
			HostIP: portConfig.Address,
			HostPort: strconv.Itoa(int(portConfig.HostInt)),
		}}
	}

	return exposedPorts, portBindings

}

// func main() {
// 	imageName := "envoy-test-app:latest"
// 	containerName := "Docker-web-server-example"

// 	envMap := map[string]int{
// 		"PORT": 80,
// 	}

// 	portConfig := PortConfig{
// 		AddressString: "127.0.0.1",
// 		HostInt: 80,
// 		ExposedInt: 80,
// 		Protocol: "tcp",
// 	}

// 	portConfigList := []PortConfig{portConfig}

// 	//

// 	envVars := MapToEnvArr(envMap)

// 	exposedPorts, portBindings := ConvertPortConfig(portConfigList)

// 	//

// 	ctx := context.Background()
// 	apiClient, err := client.New(client.FromEnv)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer apiClient.Close()

// 	res, err := apiClient.ContainerCreate(ctx, client.ContainerCreateOptions{
// 		Name: containerName,
// 		Image: imageName,
// 		Config: &container.Config{
// 			Env: envVars,
// 			ExposedPorts: exposedPorts,
// 		},
// 		HostConfig: &container.HostConfig{
// 			PortBindings: portBindings,
// 		},
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Container ID is:", res.ID)

// 	_, err = apiClient.ContainerStart(ctx, res.ID, client.ContainerStartOptions{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
	
// }
