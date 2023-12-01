package main

import (
	"fmt"
	"github.com/Hazegard/htb-prolab-cli/config"
	"github.com/Hazegard/htb-prolab-cli/prolabs"
	"github.com/Hazegard/htb-prolab-cli/utils"
	"io"
	"log"
	"strings"
)

func main() {

	conf := config.Parse()
	if !conf.Debug {
		log.SetOutput(io.Discard)
	}

	if conf.Help {
		conf.PrintHelp()
		return
	}

	err, instance := prolabs.IsProlabActive()
	if err != nil {
		fmt.Printf("Error checking active prolab: %s", err)
		return
	}
	if instance != "" {
		fmt.Printf("Already launched: %s, aborting...\n", instance)
		return
	}

	labs, err := prolabs.GetProlabs()
	lab := prolabs.SelectProLab(labs)
	err, vpns := prolabs.GetVpnProlab(lab)
	if err != nil {
		fmt.Printf("Error getting VPN prolab: %s", err)
	}
	vpn := prolabs.SelectVpn(vpns)

	err = prolabs.SetVpnProlab(lab, vpn)
	if err != nil {
		fmt.Printf("Error setting VPN prolab: %s", err)
		return
	}

	err, content := prolabs.GetVpnConf(vpn)
	if err != nil {
		fmt.Printf("Error retrieving OpenVPN file: %s", err)
		return
	}
	filename := ""
	if conf.OutVpn != "" {
		filename = conf.OutVpn
	} else {
		filename = strings.ReplaceAll(vpn.FriendlyName, " ", "_") + ".ovpn"
	}

	err, _ = utils.WriteToFile(filename, content)
	if err != nil {
		fmt.Printf("Error writing VPN config to %s: %s", filename, err)

	}
	fmt.Printf("Config OpenVPN dowloaded: %s\n", filename)
	return
}
