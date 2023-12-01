package prolabs

import (
	"fmt"
	"strconv"
)

func SelectProLab(labs []Lab) Lab {
	return Choice(labs, "Prolabs available:", func(lab Lab) string {
		return lab.Name
	})
}

func SelectVpn(vpn []VpnServer) VpnServer {
	return Choice(vpn, "VPN Available:", func(server VpnServer) string {
		return fmt.Sprintf("%s (%d users)", server.FriendlyName, server.CurrentClients)
	})
}

func Choice[T comparable](t []T, title string, display func(T) string) T {
	mapLab := map[int]T{}
	id := 1
	for _, lab := range t {
		mapLab[id] = lab
		id++
	}

	fmt.Println(title)
	i := 1
	for _ = range t {
		fmt.Printf("%d:\t%s\n", i, display(t[i-1]))
		i++
	}
	var choice string
	fmt.Printf("Choice: ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		return Choice(t, title, display)
	}
	j, err := strconv.Atoi(choice)
	if err != nil {
		fmt.Println(fmt.Errorf("Error parsing choice, try again: %s\n", err))
		return Choice(t, title, display)
	}
	c := mapLab[j]
	empty := *new(T)
	if c != empty {
		return c
	}
	fmt.Printf("No prolab available with id %d\n", j)
	return Choice(t, title, display)
}
