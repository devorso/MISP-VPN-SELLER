package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

/*
*

	{
	       "description": "",
	       "meta": {
	           "Name":"",
	           "Length Server":"",
	           "Length Location Server":"",
	           "Length Country":"",
	           "Range IP":[]

	       },
	       "uuid": "f41b3065-d1ff-4f95-87b9-2f6119a258d9",
	       "value": ""
	      }
*/
type DataVPNMeta struct {
	Name                 string `json:"name"`
	LengthServer         string `json:"length_server"`
	LengthLocationServer string `json:"length_location_server"`
	LengthCountry        string `json:"length_country"`
	IPRange              string `json:"ip_range"`
}
type DataVPN struct {
	Description string      `json:"description"`
	Uuid        string      `json:"uuid"`
	Meta        DataVPNMeta `json:"meta"`
	Value       string      `json:"value"`
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
		uuid, _ := uuid.NewUUID()

		data := DataVPN{
			Value:       Name,
			Description: Name,
			Meta: DataVPNMeta{
				Name:                 Name,
				LengthCountry:        LengthCountry,
				LengthServer:         LengthServer,
				LengthLocationServer: LengthLocationServer,
			},
			Uuid: uuid.String(),
		}
		listvpn = append(listvpn, data)
		log.Println(data)
	}

	bytes, _ := json.Marshal(&listvpn)
	err = os.WriteFile("data.json", bytes, 777)
	if err != nil {
		log.Fatal(err)
	}

}
