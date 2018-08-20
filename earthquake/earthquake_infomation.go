package earthquake

import (
	"encoding/json"
	"log"
)

type EarthquakeInformation struct {
	Result EarthquakeResult
	Report_Id string
}

type EarthquakeResult struct {
	Status   string
	Message  string
}

func ParseJsonOf(jsonText string) EarthquakeInformation {
	var info EarthquakeInformation
	jsonBytes := []byte(jsonText)
	err := json.Unmarshal(jsonBytes, &info)
	if err != nil {
		log.Fatal(err)
	}
	return info
}
