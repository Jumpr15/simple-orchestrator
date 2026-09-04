package containerTypes 

import ()

type ContainerConfig struct {
	DesiredNumContainers int
	ImageName string
	EnvMap map[string]int // change to string any later
}

type ContainerState struct {
	ID string
	State string // enum
}

type ContainerList []ContainerState