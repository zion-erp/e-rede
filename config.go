package erede

import "fmt"

type Config struct {
	ApiUrl          string
	ApiTokenization string
	Filiation       string // PV
	Sandbox         bool
	Token           string
	Version         string
}

func (c *Config) Init(filiation string, token string, version string, sandbox bool) {
	c.Filiation = filiation
	c.Token = token
	c.Version = version
	c.Sandbox = sandbox
	if sandbox {
		c.ApiUrl = fmt.Sprintf("https://api.userede.com.br/desenvolvedores/%s", version)
		c.ApiTokenization = fmt.Sprintf("https://rl7-sandbox-api.useredecloud.com.br/token-service/%s", version)
	} else {
		c.ApiUrl = fmt.Sprintf("https://api.userede.com.br/erede/%s", version)
		c.ApiTokenization = fmt.Sprintf("https://api.userede.com.br/redelabs/token-service/%s", version)
	}
}
