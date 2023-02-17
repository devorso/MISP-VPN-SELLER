package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type DataVPNMeta struct {
	Name                 string   `json:"name"`
	LengthServer         string   `json:"length_server"`
	LengthLocationServer string   `json:"length_location_server"`
	LengthCountry        string   `json:"length_country"`
	IPRange              []string `json:"ip_range"`
}

type DataVPN struct {
	Description string      `json:"description"`
	Uuid        string      `json:"uuid"`
	Meta        DataVPNMeta `json:"meta"`
	Value       string      `json:"value"`
}

type Source struct {
	Link  string
	Index uint
}

type DataIPMeta struct {
	IP        string `json:"ip"`
	Date      string `json:"date"`
	VPNSeller string `json:"vpn_seller"`
}
type DataIP struct {
	Description string     `json:"description"`
	Uuid        string     `json:"uuid"`
	Meta        DataIPMeta `json:"meta"`
	Value       string     `json:"value"`
}

func main() {
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	start := 5
	end := 24
	//A to E
	var listvpn []DataVPN
	for i := start; i < end; i++ {
		Name, _ := f.GetCellValue("Feuil1", fmt.Sprintf("%v%v", "A", i))
		LengthServer, _ := f.GetCellValue("Feuil1", fmt.Sprintf("%v%v", "B", i))
		LengthLocationServer, _ := f.GetCellValue("Feuil1", fmt.Sprintf("%v%v", "C", i))
		LengthCountry, _ := f.GetCellValue("Feuil1", fmt.Sprintf("%v%v", "D", i))
		uuidValue, _ := uuid.NewUUID()
		var listIp []string
		data := DataVPN{
			Value:       Name,
			Description: Name,
			Meta: DataVPNMeta{
				Name:                 Name,
				LengthCountry:        LengthCountry,
				LengthServer:         LengthServer,
				LengthLocationServer: LengthLocationServer,
				IPRange:              listIp,
			},
			Uuid: uuidValue.String(),
		}
		listvpn = append(listvpn, data)
		log.Println(data)
	}
	//Now Parsing ip.
	var listIp []DataIP
	///var listSource []Source

	/**
	https://raw.githubusercontent.com/silvether/pia-vpn-ranges/master/IPs_022019
	https://gist.githubusercontent.com/triggex/c6bc554410a84ea1b3ef1c19c5a92d49/raw/1d5a60401f631356d156c21b32471d15dff2e0e1/NordVPN-Server-IP-List-2020.txt
	https://github.com/Luen/IPVanish-Server-List
	https://www.betamaster.us/blog/?p=561
	https://github.com/scriptzteam/ProtonVPN-VPN-Ips
	https://mullvad.net/fr/servers/
	https://support.vyprvpn.com/hc/en-us/articles/360037728912-What-are-the-VyprVPN-server-addresses-
	https://gist.githubusercontent.com/JamoCA/eedaf4f7cce1cb0aeb5c1039af35f0b7/raw/01faaa9ffa025a66be0abfd674bf111090d9ebb9/NordVPN-Server-IP-List.txt
	*/
	// NordVPN ip

	linksNordVPN := []string{"https://gist.githubusercontent.com/triggex/c6bc554410a84ea1b3ef1c19c5a92d49/raw/1d5a60401f631356d156c21b32471d15dff2e0e1/NordVPN-Server-IP-List-2020.txt", "https://gist.githubusercontent.com/JamoCA/eedaf4f7cce1cb0aeb5c1039af35f0b7/raw/01faaa9ffa025a66be0abfd674bf111090d9ebb9/NordVPN-Server-IP-List.txt"}
	dateNordVPN := []string{"Fev 04, 2020", "Jan 09, 2020"}
	for i, v := range linksNordVPN {
		resp, err := http.Get(v)
		if err != nil {
			log.Println(err)
		}

		body, err := io.ReadAll(resp.Body)
		ipList := strings.Split(string(body), "\n")
		uuidValIp, _ := uuid.NewUUID()
		for _, ip := range ipList {
			if len(ip) > 0 {
				dataIp := DataIP{
					Value:       ip,
					Description: "IP Used in NORDVPN",
					Uuid:        uuidValIp.String(),
					Meta: DataIPMeta{
						IP:        ip,
						Date:      dateNordVPN[i],
						VPNSeller: "NordVPN",
					},
				}
				listIp = append(listIp, dataIp)
			}
		}
	}
	//Private access vpn list
	resp, err := http.Get("https://raw.githubusercontent.com/silvether/pia-vpn-ranges/master/IPs_022019")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	IPListPIA := strings.Split(string(body), "\n")

	for d := 6; d < len(IPListPIA); d++ {
		uuidValIp, _ := uuid.NewUUID()
		ipD := IPListPIA[d]
		if len(ipD) > 0 {
			ipdataPia := DataIP{
				Value: ipD,
				Uuid:  uuidValIp.String(),
				Meta: DataIPMeta{
					IP:        ipD,
					Date:      "Dec 1, 2019",
					VPNSeller: "Pia",
				},
				Description: "IP Used for Private Access VPN",
			}
			listIp = append(listIp, ipdataPia)
		}
	}

	resp, err = http.Get("https://raw.githubusercontent.com/Luen/IPVanish-Server-List/master/ipvanish-allowlist.txt")
	if err != nil {
		log.Fatal(err)
	}
	body, err = io.ReadAll(resp.Body)
	IPListVanish := strings.Split(string(body), "\n")
	for k := 3; k < len(IPListVanish); k++ {
		ipData := strings.Split(IPListVanish[k], ":")
		if len(ipData) > 1 {
			ipDataValue := ipData[1]
			uuidValIp, _ := uuid.NewUUID()

			if len(ipData) > 0 {
				ipdataPia := DataIP{
					Value: ipDataValue,
					Uuid:  uuidValIp.String(),
					Meta: DataIPMeta{
						IP:        ipDataValue,
						Date:      "Aug 26, 2015",
						VPNSeller: "IPVanish",
					},
					Description: "IP Used for IPVanish",
				}
				listIp = append(listIp, ipdataPia)
			}
		}
	}
	var protonVPNLinks []string = []string{"https://raw.githubusercontent.com/scriptzteam/ProtonVPN-VPN-IPs/main/entry_ips.txt", "https://raw.githubusercontent.com/scriptzteam/ProtonVPN-VPN-IPs/main/exit_ips.txt"}
	var protonVPNDate []string = []string{"Jul 20, 2022", "Jul 14, 2022"}
	for i, v := range protonVPNLinks {
		// 0 entry
		// 1 exit
		respData, err := http.Get(v)
		if err != nil {
			log.Fatal(err)
		}
		body, err = io.ReadAll(respData.Body)
		IPListProtonIp := strings.Split(string(body), "\n")
		for _, ip := range IPListProtonIp {
			if len(ip) > 0 {
				description := "Entry IP for Proton VPN"
				if i > 0 {
					description = "Exit IP for Proton VPN"
				}
				uuidValIp, _ := uuid.NewUUID()
				dataI := DataIP{Value: ip, Description: description, Uuid: uuidValIp.String(), Meta: DataIPMeta{
					IP:        ip,
					Date:      protonVPNDate[i],
					VPNSeller: "ProtonVPN",
				}}
				listIp = append(listIp, dataI)
			}
		}

	}
	//Mullvad
	/*https://mullvad.net/fr/servers/
	c := colly.NewCollector()

	// Find and visit all links

	c.OnError(func(_ *colly.Response, err error) {
		log.Println(err.Error())
	})
	c.OnRequest(func(request *colly.Request) {
		log.Println(request.Headers)
	})
	c.OnHTML("[aria-controls=\"b6541d73-8ccd-490f-b366-87490b32cb17\"]", func(e *colly.HTMLElement) {

		fmt.Println(e)
	})
	c.Visit("https://mullvad.net/fr/servers/")
	*/
	bytes, _ := json.Marshal(&listIp)
	err = os.WriteFile("ip_galaxy.json", bytes, 777)
	if err != nil {
		//log.Fatal(err)
	}
	bytesD, _ := json.Marshal(&listvpn)
	err = os.WriteFile("data.json", bytesD, 777)
	if err != nil {
		//log.Fatal(err)
	}

}
