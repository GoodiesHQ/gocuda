package cuda

// Network Objects - IPs, references, hostnames, networks
type NetworkObjectEntry struct {
	IP        string `json:"ip,omitempty"`
	MAC       string `json:"mac,omitempty"`
	Interface string `json:"interface,omitempty"`
	Comment   string `json:"comment,omitempty"`
}

type NetworkObjectExcludedEntry struct {
	Entry *NetworkObjectEntry `json:"entry"`
}

type NetworkObjectExcludedEntryUpdate struct {
	Entry *NetworkObjectEntry `json:"entry"`
	ID    string              `json:"id"`
}

type NetworkObjectIncludedEntry struct {
	Entry      *NetworkObjectEntry `json:"entry,omitempty"`
	References string              `json:"references,omitempty"`
}

type NetworkObjectIncludedEntryUpdate struct {
	Entry      *NetworkObjectEntry `json:"entry,omitempty"`
	References string              `json:"references,omitempty"`
	ID         string              `json:"id"`
}

type NetworkObjectGeoEntry struct {
	Included []string `json:"included,omitempty"`
	Excluded []string `json:"excluded,omitempty"`
}

type NetworkObjectGeoEntryChange struct {
	Included []string `json:"included,omitempty"`
	Excluded []string `json:"excluded,omitempty"`
}

type NetworkObjectType string

const (
	NetworkObjectTypeGeneric       NetworkObjectType = "generic"
	NetworkObjectTypeFQDN          NetworkObjectType = "hostname"
	NetworkObjectTypeHostname      NetworkObjectType = "hostname"
	NetworkObjectTypeIPv4Address   NetworkObjectType = "singleIPv4Address"
	NetworkObjectTypeIPv6Address   NetworkObjectType = "singleIPv6Address"
	NetworkObjectTypeIPv4Addresses NetworkObjectType = "listIPv4Address"
	NetworkObjectTypeIPv6Addresses NetworkObjectType = "listIPv6Address"
	NetworkObjectTypeIPv4Network   NetworkObjectType = "singleIPv4Network"
	NetworkObjectTypeIPv6Network   NetworkObjectType = "singleIPv6Network"
	NetworkObjectTypeIPv4Networks  NetworkObjectType = "listIPv4Network"
	NetworkObjectTypeIPv6Networks  NetworkObjectType = "listIPv6Network"
)

// The complete network object used by Barracuda for firewall and application rules
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

type NetworkObjectGeoChange struct {
	Add    []string `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
	Update []string `json:"update,omitempty"`
}

type NetworkObjectGeoUpdate struct {
	Included []NetworkObjectGeoChange `json:"included,omitempty"`
	Excluded []NetworkObjectGeoChange `json:"excluded,omitempty"`
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
	Geo         NetworkObjectGeoUpdate      `json:"geo,omitempty"`
}
