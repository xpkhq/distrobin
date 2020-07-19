// Package configure provides configure
// prototypes of distrobin implemention
package configure

// ClientConfigure consists of the required options
// for the client
type ClientConfigure struct {
	// Pool is the address of Pool service
	Pool string
	// Gzip determines enables gzip or not
	Gzip bool
}
