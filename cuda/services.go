package cuda

// Services - services running on the firewall like eventing, firewall, dhcp

type Service struct {
	Name            string `json:"name"`
	ModuleName      string `json:"moduleName"`
	State           string `json:"state"`
	Processes       int    `json:"processes"`
	FileDescriptors int    `json:"fileDescriptors"`
	Memory          int    `json:"memory"`
	Info            string `json:"info"`
}

type Services struct {
	Services []Service `json:"services"`
}

type ServicesSimple struct {
	Services []string `json:"services"`
}
