package scheduler

import (
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
func (sch *Scheduler) ScheduleNode() { 
	
}

func (sch *Scheduler) GetState() {
	log.Println("Scheduler up and running!")
}