package cuda

// Sessions: Actives sessions within the firewall
type Session struct {
	Admin string `json:"admin"`
	Idle  int    `json:"idle"`
	Name  string `json:"name"`
	Peer  string `json:"peer"`
	PID   int    `json:"pid"`
	Type  string `json:"type"`
}

type Sessions struct {
	Sessions []Session `json:"sessions"`
}
