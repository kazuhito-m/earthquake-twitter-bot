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

func (gc GeoCode) Ok() bool {
	return gc.Status == "OK"
}

func (gc GeoCode) MajorPlacesIncludingPrefectureName() []string {
	texts := []string{}
	texts = append(texts, gc.Plus_Code.Compound_Code)
	for _, result := range gc.Results {
		for _, addressComponent := range result.Address_Components {
			if contains(addressComponent.Types, "administrative_area_level_1") {
				texts = append(texts, addressComponent.Long_Name)
			}
		}
		texts = append(texts, result.Formatted_Address)
	}
	return texts
}

func contains(texts []string, text string) bool {
	for _, v := range texts {
		if text == v {
			return true
		}
	}
	return false
}
