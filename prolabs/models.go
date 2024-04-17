package prolabs

import (
	"fmt"
	"time"
)

type Displayer interface {
	Display() string
}

type Prolabs struct {
	Status bool `json:"status"`
	Data   struct {
		Count int   `json:"count"`
		Labs  []Lab `json:"labs"`
	} `json:"data"`
}

type Lab struct {
	Id                         int         `json:"id"`
	Name                       string      `json:"name"`
	ReleaseAt                  time.Time   `json:"release_at"`
	ProMachinesCount           int         `json:"pro_machines_count"`
	ProFlagsCount              int         `json:"pro_flags_count"`
	Ownership                  int         `json:"ownership"`
	UserEligibleForCertificate bool        `json:"user_eligible_for_certificate"`
	New                        bool        `json:"new"`
	SkillLevel                 string      `json:"skill_level"`
	DesignatedCategory         string      `json:"designated_category"`
	Team                       string      `json:"team"`
	Level                      int         `json:"level"`
	LabServersCount            int         `json:"lab_servers_count"`
	CoverImgUrl                interface{} `json:"cover_img_url"`
}

func (l Lab) Display() string {
	return l.Name
}

type VpnData struct {
	Status bool `json:"status"`
	Data   struct {
		Assigned interface{} `json:"assigned"`
		Options  struct {
			EU struct {
				EUProLab struct {
					Location string            `json:"location"`
					Name     string            `json:"name"`
					Servers  map[int]VpnServer `json:"servers"`
				} `json:"EU - Pro Lab"`
			} `json:"EU"`
			US struct {
				USProLab struct {
					Location string            `json:"location"`
					Name     string            `json:"name"`
					Servers  map[int]VpnServer `json:"servers"`
				} `json:"US - Pro Lab"`
			} `json:"US"`
		} `json:"options"`
	} `json:"data"`
}

type VpnServer struct {
	Id             int    `json:"id"`
	FriendlyName   string `json:"friendly_name"`
	Full           bool   `json:"full"`
	CurrentClients int    `json:"current_clients"`
	Location       string `json:"location"`
}

func (v VpnServer) Display() string {
	return fmt.Sprintf("%s (%d users)", v.FriendlyName, v.CurrentClients)
}

type SwitchVpn struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Id             int    `json:"id"`
		FriendlyName   string `json:"friendly_name"`
		CurrentClients int    `json:"current_clients"`
		Location       string `json:"location"`
	} `json:"data"`
}

type ProlabConnection struct {
	CanAccess            bool   `json:"can_access"`
	LocationTypeFriendly string `json:"location_type_friendly"`
	AssignedServer       struct {
		Id             int    `json:"id"`
		FriendlyName   string `json:"friendly_name"`
		CurrentClients int    `json:"current_clients"`
		Location       string `json:"location"`
	} `json:"assigned_server"`
}

func (plc *ProlabConnection) IsConnected() bool {
	return !(plc.LocationTypeFriendly == "" && plc.AssignedServer.Id == 0)
}

type CurrentConnections struct {
	Data struct {
		Lab struct {
			CanAccess            bool   `json:"can_access"`
			LocationTypeFriendly string `json:"location_type_friendly"`
			AssignedServer       struct {
				Id             int    `json:"id"`
				FriendlyName   string `json:"friendly_name"`
				CurrentClients int    `json:"current_clients"`
				Location       string `json:"location"`
			} `json:"assigned_server"`
		} `json:"lab"`
		StartingPoint struct {
			CanAccess            bool   `json:"can_access"`
			LocationTypeFriendly string `json:"location_type_friendly"`
			AssignedServer       struct {
				Id             int    `json:"id"`
				FriendlyName   string `json:"friendly_name"`
				CurrentClients int    `json:"current_clients"`
				Location       string `json:"location"`
			} `json:"assigned_server"`
		} `json:"starting_point"`
		Endgames struct {
			CanAccess            bool   `json:"can_access"`
			LocationTypeFriendly string `json:"location_type_friendly"`
			AssignedServer       struct {
				Id             int    `json:"id"`
				FriendlyName   string `json:"friendly_name"`
				CurrentClients int    `json:"current_clients"`
				Location       string `json:"location"`
			} `json:"assigned_server"`
		} `json:"endgames"`
		Fortresses struct {
			CanAccess      bool        `json:"can_access"`
			AssignedServer interface{} `json:"assigned_server"`
		} `json:"fortresses"`
		ProLabs     map[string]any `json:"pro_labs"`
		Competitive struct {
			CanAccess      bool `json:"can_access"`
			AssignedServer struct {
				Id             int    `json:"id"`
				FriendlyName   string `json:"friendly_name"`
				CurrentClients int    `json:"current_clients"`
				Location       string `json:"location"`
			} `json:"assigned_server"`
			Available            bool   `json:"available"`
			LocationTypeFriendly string `json:"location_type_friendly"`
			Machine              struct {
				Id             int    `json:"id"`
				Name           string `json:"name"`
				AvatarThumbUrl string `json:"avatar_thumb_url"`
			} `json:"machine"`
		} `json:"competitive"`
	} `json:"data"`
	Status bool `json:"status"`
}

type Status struct {
	Type                 string `json:"type"`
	LocationTypeFriendly string `json:"location_type_friendly"`
	Server               struct {
		Id           int    `json:"id"`
		Hostname     string `json:"hostname"`
		Port         int    `json:"port"`
		FriendlyName string `json:"friendly_name"`
	} `json:"server"`
	Connection struct {
		Name          string  `json:"name"`
		ThroughPwnbox bool    `json:"through_pwnbox"`
		Ip4           string  `json:"ip4"`
		Ip6           string  `json:"ip6"`
		Down          float64 `json:"down"`
		Up            float64 `json:"up"`
	} `json:"connection"`
}
