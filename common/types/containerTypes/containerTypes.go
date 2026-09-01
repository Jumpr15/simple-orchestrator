package containerTypes 

type ContainerConfig struct {
	
}

type ContainerState struct {
	ID string
	State string // enum
}

type ContainerList []ContainerState