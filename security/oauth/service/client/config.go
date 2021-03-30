package client

import (
	"net/url"
	"time"

	"golang.org/x/oauth2/clientcredentials"
)

type Endpoint struct {
	TokenURL  	 string    `yaml:"tokenURL" json:"tokenURL" envconfig:"TOKEN_URL"`
}

type Config struct {
	ClientID       string        `yaml:"clientID" json:"clientID" envconfig:"CLIENT_ID"`
	ClientSecret   string        `yaml:"clientSecret" json:"clientSecret" envconfig:"CLIENT_SECRET"`
	Scopes         []string      `yaml:"scopes" json:"scopes" envconfig:"SCOPES"`
	Endpoint                     `yaml:",inline"`
	Audience       string        `yaml:"audience" json:"audience" envconfig:"AUDIENCE"`
	RequestTimeout time.Duration `yaml:"timeout" json:"timeout" envconfig:"TIMEOUT" default:"10s"`
}

// ToClientCrendtials converts to clientcredentials.Config
func (c Config) ToClientCrendtials() clientcredentials.Config {
	v := make(url.Values)
	if c.Audience != "" {
		v.Set("audience", c.Audience)
	}

	return clientcredentials.Config{
		ClientID:       c.ClientID,
		ClientSecret:   c.ClientSecret,
		Scopes:         c.Scopes,
		TokenURL:       c.Endpoint.TokenURL,
		EndpointParams: v,
	}
}
