package domain

type Param struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled"`
}

type Header struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled"`
}

type BasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type BearerToken struct {
	Token string `json:"token"`
}
type OAuth20 struct {
	Key              string `json:"key"`              // key
	OidcDiscoveryURL string `json:"oidcDiscoveryURL"` // OpenID Connect Discovery URL
	AuthURL          string `json:"authURL"`          // Authentication URL
	AccessTokenURL   string `json:"accessTokenURL"`   // Access Token URL
	ClientID         string `json:"clientID"`         // Client ID
	Scope            string `json:"scope"`            // Scope
}
type ApiKey struct {
	Key          string `json:"key"`
	Value        string `json:"value"`
	TransferMode string `json:"transferMode"`
}
