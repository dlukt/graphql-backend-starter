package config

type OIDCConfig struct {
	Issuer          string
	Audience        string
	AuthorizedParty string
}

var OidcConfigDev = OIDCConfig{
	Issuer:          "https://auth.icod.de/realms/dev",
	Audience:        "spa",
	AuthorizedParty: "spa",
}

var OidcConfigProd = OIDCConfig{
	Issuer:          "https://auth.icod.de/realms/dev",
	Audience:        "spa",
	AuthorizedParty: "spa",
}
