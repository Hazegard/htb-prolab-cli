package config

import (
	"fmt"
	"github.com/spf13/pflag"
)

type Config struct {
	Debug  bool
	OutVpn string
	Help   bool
}

func Parse() *Config {
	debug := false
	help := false
	vpnfile := ""
	pflag.BoolVarP(&debug, "debug", "d", false, "debug mode")
	pflag.BoolVarP(&help, "help", "h", false, "print help")
	pflag.StringVarP(&vpnfile, "output", "o", "", "Target of the downloaded OpenVPN file")
	pflag.Parse()

	return &Config{
		Debug:  debug,
		OutVpn: vpnfile,
		Help:   help,
	}
}

func (c *Config) PrintHelp() {
	fmt.Println("Usage of htb-prolab-cli:")
	pflag.PrintDefaults()
}
