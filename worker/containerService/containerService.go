package containerService

import (
	ctrType "simple-orchestrator/common/types/containerTypes"

	"log"
)

func CreateAndStartContainer(kv *ristretto.Cache[string, any], config ctrType.ContainerConfig) {
	log.Println("createandstartcontainer")

	// inc container count
	currentNumContainersInterface, found := kv.Get("currentNumContainers")
	if !found {
		log.Fatal("currentNumContainers not found in kv")
	}
	currentNumContainers, ok := currentNumContainersInterface.(int)
	if !ok {
		log.Fatal("currentNumContainers is not an int?")
	}
	// insert container info into kv
	kv.Set("currentNumContainers", currentNumContainers+1, 1)
	// append container to list of containers for node in kv
	

}