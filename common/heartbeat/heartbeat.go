package heartbeat

import ()

type ContainerConfig struct {

}

type ContainerState struct {
	ID string
	State string // enum
	
}

type ContainerList []ContainerState

type NodeAddrConfig struct {
	Id string
	Address string
	Port string
}

type Args struct {
	NodeAddrConfig
	ContainerConfig

	DesiredNumContainers int // Needed? (Running or Total )
}

type Response struct {
	NodeAddrConfig
	ContainerList 

	CurrentNumContainers int // Needed?
}

type HeartbeatServer struct {

}

func (hb *HeartbeatServer) SendHealthcheck(args *Args, res *Response) error {
	res = &Response{}
	return nil
}