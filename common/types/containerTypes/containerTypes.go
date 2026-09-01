package containerTypes 

type ContainerConfig struct {
	DesiredNumContainers int
}

type ContainerState struct {
	ID string
	State string // enum
}

type ContainerList []ContainerState