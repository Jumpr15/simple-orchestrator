package scheduler

import (
	ctrType "simple-orchestrator/common/types/containerTypes"
	"simple-orchestrator/master/kvStore"

	"log"
)

// maybe turn into a interface to allow for different scheduling algos

type Scheduler struct {
	kvStore *kvStore.KvStore
	State string 
}

func New(kv *kvStore.KvStore) *Scheduler {
	sch := Scheduler{kvStore: kv, State: "scheduler state!"}
	return &sch
}

// basic scheduler, tells nodes to spin up 2 instances of a container
func (sch *Scheduler) ScheduleNode() ctrType.ContainerConfig { 
	// replace with call to ctr cfg constructor later 
	
	cfg := ctrType.ContainerConfig{
		DesiredNumContainers: 1,
		ImageName: "envoy-test-app:latest",
		EnvMap: map[string]int{
			"PORT": 80,
		},
		PortConfig: ctrType.PortConfig{
			AddressString: "127.0.0.1",
			HostInt: 80,
			ExposedInt: 80,
			Protocol: "tcp",			
		},
	}
	return cfg
}

func (sch *Scheduler) GetState() {
	log.Println("Scheduler up and running!")
}