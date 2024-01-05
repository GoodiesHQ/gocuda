package cuda

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
