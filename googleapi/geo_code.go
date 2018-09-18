package googleapi

import "encoding/json"

type GeoCode struct {
	Plus_Code struct {
		Compound_Code string
	}
	Results []  struct {
		Address_Components [] struct {
			Long_Name string
			Types     []string
		}
		Formatted_Address string
	}
	Status string
}

func ParseJsonOf(jsonText string) GeoCode {
	var geoCode GeoCode
	jsonBytes := []byte(jsonText)
	json.Unmarshal(jsonBytes, &geoCode)
	return geoCode
}
