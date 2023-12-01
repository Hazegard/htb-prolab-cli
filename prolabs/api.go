package prolabs

import (
	"encoding/json"
	"fmt"
	"github.com/GoToolSharing/htb-cli/lib/utils"
	"io"
	"net/http"
	"sort"
)

const (
	BaseHackTheBoxAPIURL = "https://www.hackthebox.com/api/v4"
	PROLABS_ROUTE        = "prolabs"
)

func GetProlabs() ([]Lab, error) {

	url := BaseHackTheBoxAPIURL + "/" + PROLABS_ROUTE
	resp, err := utils.HtbRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to download prolabs list: %s", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to download prolabs list: %s", err)
	}

	var data Prolabs
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("unable to download prolabs list: %s", err)
	}
	labs := data.Data.Labs
	sort.Slice(labs, func(i, j int) bool {
		return labs[i].Name < labs[j].Name
	})
	return data.Data.Labs, nil
}

func GetVpnProlab(lab Lab) (error, []VpnServer) {
	url := BaseHackTheBoxAPIURL + fmt.Sprintf("/connections/servers/prolab/%d", lab.Id)
	resp, err := utils.HtbRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("unable to download vpn list for prolab %s list: %s", lab.Name, err), nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error getting VPN prolab: %s", err), nil
	}

	var data VpnData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return fmt.Errorf("error getting VPN prolab: %s", err), nil
	}
	vpn := data.Data.Options.EU.EUProLab.Servers

	servers := []VpnServer{}
	for _, server := range vpn {
		servers = append(servers, server)
	}

	sort.Slice(servers, func(i, j int) bool {
		return servers[i].FriendlyName < servers[j].FriendlyName
	})
	return nil, servers
}

func SetVpnProlab(lab Lab, vpn VpnServer) error {
	url := BaseHackTheBoxAPIURL + fmt.Sprintf("/connections/servers/switch/%d", vpn.Id)
	req_body := "{\"arena\":false}"
	resp, err := utils.HtbRequest(http.MethodPost, url, []byte(req_body))
	if err != nil {
		return fmt.Errorf("unable set vpn to %s (Prolab: %s)", vpn.FriendlyName, lab.Name)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable set vpn to %s (Prolab: %s)", vpn.FriendlyName, lab.Name)

	}

	var response SwitchVpn
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("unable set vpn to %s (Prolab: %s)", vpn.FriendlyName, lab.Name)

	}
	if !response.Status {
		return fmt.Errorf("error setting the vpn: %s", response.Message)
	}
	return nil
}

func GetCurrentProlabConnections() (error, []ProlabConnection) {
	url := BaseHackTheBoxAPIURL + "/connections"
	resp, err := utils.HtbRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("error getting vpn connection list: %s", err), nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	var currentProlabConnection CurrentConnections
	err = json.Unmarshal(body, &currentProlabConnection)
	if err != nil {
		return err, nil
	}

	var plc []ProlabConnection
	for _, pl := range currentProlabConnection.Data.ProLabs {
		jsonString, _ := json.Marshal(pl)

		// convert json to struct
		s := ProlabConnection{}
		err := json.Unmarshal(jsonString, &s)
		if err != nil {
			continue
		}
		plc = append(plc, s)
	}

	return nil, plc
}

func IsProlabActive() (error, string) {
	url := BaseHackTheBoxAPIURL + "/connection/status"
	resp, err := utils.HtbRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("error getting connection status: %s", err), ""
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, ""
	}

	var status []Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		return err, ""
	}

	for _, statu := range status {
		if statu.Type == "Pro Lab" {
			return nil, statu.Server.FriendlyName
		}
	}
	return nil, ""
}

func GetVpnConf(vpn VpnServer) (error, []byte) {
	url := fmt.Sprintf("%s/access/ovpnfile/%d/0", BaseHackTheBoxAPIURL, vpn.Id)
	resp, err := utils.HtbRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("error getting openvpn profile: %s", err), nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %s"), nil
	}

	return nil, body
}
