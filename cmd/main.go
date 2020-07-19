// Command distrobin provides interfaces
// to distrobin
package main

import (
	"os"

	"github.com/urfave/cli/v2"
	"github.com/xpkhq/distrobin"
)

func main() {
	app := &cli.App{
		Name:      "distrobin",
		Usage:     "DistroBin's official cli utility",
		UsageText: "distrobin command [options] [arguments]",
		Version:   distrobin.Version,
	}
	app.Run(os.Args)
}
