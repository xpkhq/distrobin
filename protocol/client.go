package protocol

import "github.com/xpkhq/distrobin/configure"

// Client is the client class
type Client struct {
	conf *configure.ClientConfigure
}

// NewClient creates a new client
func NewClient(conf *configure.ClientConfigure) *Client {
	return &Client{
		conf: conf,
	}
}
