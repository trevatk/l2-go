package credential

import (
	"encoding/base64"
)

// ICredentials
type ICredentials interface {
	Key() string
	Secret() string
}

// Credentials
type Credentials struct {
	key    string
	secret string
}

// interface verification
var _ ICredentials = (*Credentials)(nil)

// Config
type Config struct {
	Key    string
	Secret string
}

// NewCredentials
func NewFromConfig(cfg *Config) *Credentials {

	key := base64.StdEncoding.EncodeToString([]byte(cfg.Key))

	secret := base64.StdEncoding.EncodeToString([]byte(cfg.Secret))

	return &Credentials{key: key, secret: secret}
}

// Key
func (c *Credentials) Key() string {
	return c.key
}

// Secret
func (c *Credentials) Secret() string {
	return c.secret
}
