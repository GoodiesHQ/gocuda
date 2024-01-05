package cuda

// Network Objects - IPs, references, hostnames, networks
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

type NetworkObjectGeoEntryChange struct {
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
