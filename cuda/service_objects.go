package cuda

type ServiceObject struct {
	Color string `json:"color,omitempty"`
}

// Network Objects - IPs, references, hostnames, networks
type ServiceObjectEntry struct {
	BalancedTimeout    int                                  `json:"balancedTimeout,omitempty"`
	Comment            string                               `json:"comment,omitempty"`
	Plugin             string                               `json:"plugin,omitempty"`
	Protocol           string                               `json:"protocol"`
	ICMP               *ServiceObjectEntryICMP              `json:"icmp,omitempty"`
	ProtocolProtection ServiceObjectEntryProtocolProtection `json:"protocolProtection,omitempty"`
	SessionTimeout     int                                  `json:"sessionTimeout,omitempty"`
	TCP                *ServiceObjectEntryPorts             `json:"tcp,omitempty"`
	UDP                *ServiceObjectEntryPorts             `json:"udp,omitempty"`
}

type ServiceObjectEntryPortRange struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type ServiceObjectEntryPorts struct {
	ClientPortsUsed ServiceObjectEntryPortRange `json:"clientPortsUsed,omitempty"`
	DynamicService  string                      `json:"dynamicService,omitempty"`
	Ports           []string                    `json:"ports"`
	ServiceLabel    string                      `json:"serviceLabel,omitempty"`
}

type ServiceObjectEntryICMP struct {
	MaxSize  int `json:"maxSize"`
	MinDelay int `json:"minDelay"`
}

const (
	// protocol protection actions
	ServiceObjectEntryProtocolProtectionActionNone   = "none"
	ServiceObjectEntryProtocolProtectionActionReport = "report"
	ServiceObjectEntryProtocolProtectionActionReset  = "reset"
	ServiceObjectEntryProtocolProtectionActionDrop   = "drop"

	// protocol protection policies
	ServiceObjectEntryProtocolProtectionPolicyWhitelist = "whitelist"
	ServiceObjectEntryProtocolProtectionPolicyBlacklist = "blacklist"
)

type ServiceObjectEntryProtocolProtection struct {
	Action    string   `json:"action"`
	Policy    string   `json:"policy"`
	Protocols []string `json:"protocols"`
}
