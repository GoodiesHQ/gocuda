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

// Admins: Box Administrators
type Admin struct {
	Name                  string   `json:"name"`
	FullName              string   `json:"fullName"`
	Enabled               bool     `json:"enabled"`
	Roles                 []string `json:"roles"`
	AuthenticationLevel   string   `json:"authenticationLevel"`
	ExternalLoginName     string   `json:"externalLoginName"`
	EnforcePasswordChange bool     `json:"enforcePasswordChange"`
	GracePeriod           int      `json:"gracePeriod"`
	LoginEvent            string   `json:"loginEvent"`
	NextForcedChange      int      `json:"nextForcedChange"`
	PasswordChangeMode    string   `json:"passwordChangeMode"`
	PasswordValidation    string   `json:"passwordValidation"`
	PeerIPRestriction     []string `json:"peerIpRestriction"`
	PublicKey             string   `json:"publicKey"`
	SystemLevelAccess     string   `json:"systemLevelAccess"`
	WarningPeriod         int      `json:"warningPeriod"`
}

type Admins struct {
	Admins []Admin `json:"admins"`
}

type AdminsSimple struct {
	Admins []string `json:"admins"` /* user names only */
}

/* Services */

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

/* Network Objects */
type NetworkObjectEntry struct {
	Description string `json:"description,omitempty"`
	IP          string `json:"ip,omitempty"`
	MAC         string `json:"mac,omitempty"`
	Interface   string `json:"interface,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

type NetworkObjectExcludedEntry struct {
	Entry NetworkObjectEntry `json:"entry"`
}

type NetworkObjectExcludedEntryUpdate struct {
	Entry NetworkObjectEntry `json:"entry"`
	ID    string             `json:"id"`
}

type NetworkObjectIncludedEntry struct {
	Entry      NetworkObjectEntry `json:"entry,omitempty"`
	References string             `json:"references,omitempty"`
}

type NetworkObjectIncludedEntryUpdate struct {
	Entry      NetworkObjectEntry `json:"entry,omitempty"`
	References string             `json:"references,omitempty"`
	ID         string             `json:"id"`
}

type NetworkObjectGeoEntry struct {
	Included []string `json:"included,omitempty"`
	Excluded []string `json:"excluded,omitempty"`
}

type NetworkObjectType string

const NetworkObjectTypeGeneric = "generic"
const NetworkObjectTypeFQDN = "hostname"
const NetworkObjectTypeHostname = "hostname"
const NetworkObjectTypeIPv4Address = "singleIPv4Address"
const NetworkObjectTypeIPv6Address = "singleIPv6Address"
const NetworkObjectTypeIPv4Addresses = "listIPv4Address"
const NetworkObjectTypeIPv6Addresses = "listIPv6Address"
const NetworkObjectTypeIPv4Network = "singleIPv4Network"
const NetworkObjectTypeIPv6Network = "singleIPv6Network"
const NetworkObjectTypeIPv4Networks = "listIPv4Network"
const NetworkObjectTypeIPv6Networks = "listIPv6Network"

type NetworkObject struct {
	Name        string                       `json:"name"`
	Type        NetworkObjectType            `json:"type"`
	Shared      bool                         `json:"shared"`
	Color       string                       `json:"color,omitempty"`
	Comment     string                       `json:"comment,omitempty"`
	DNSLifetime int                          `json:"dnsLifetime,omitempty"`
	Dynamic     bool                         `json:"dynamic"`
	Included    []NetworkObjectIncludedEntry `json:"included,omitempty"`
	Excluded    []NetworkObjectExcludedEntry `json:"excluded,omitempty"`
	Geo         *NetworkObjectGeoEntry       `json:"geo,omitempty"`
}

type NetworkObjects struct {
	NetworkObjects []NetworkObject `json:"objects"`
}

type NetworkObjectExcludedChange struct {
	Add    []NetworkObjectExcludedEntry       `json:"add,omitempty"`
	Remove []string                           `json:"remove,omitempty"`
	Update []NetworkObjectExcludedEntryUpdate `json:"update,omitempty"`
}

type NetworkObjectIncludedChange struct {
	Add    []NetworkObjectIncludedEntry       `json:"add,omitempty"`
	Remove []string                           `json:"remove,omitempty"`
	Update []NetworkObjectIncludedEntryUpdate `json:"update,omitempty"`
}

type NetworkObjectsSimple struct {
	NetworkObjects []string `json:"objects"` /* object names only */
}

type NetworkObjectUpdate struct {
	Name        string                      `json:"name"`
	Type        NetworkObjectType           `json:"type,omitempty"`
	Shared      bool                        `json:"shared"`
	Color       string                      `json:"color,omitempty"`
	Comment     string                      `json:"comment,omitempty"`
	DNSLifetime int                         `json:"dnsLifetime,omitempty"`
	Dynamic     bool                        `json:"dynamic"`
	Included    NetworkObjectIncludedChange `json:"included,omitempty"`
	Excluded    NetworkObjectExcludedChange `json:"excluded,omitempty"`
	Geo         *NetworkObjectGeoEntry      `json:"geo,omitempty"`
}
