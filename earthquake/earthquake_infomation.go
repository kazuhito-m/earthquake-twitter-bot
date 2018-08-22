package earthquake

import (
	"encoding/json"
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
	json.Unmarshal(jsonBytes, &info)
	return info
}
