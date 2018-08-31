package earthquake

import (
	"encoding/json"
)

type EarthquakeReport struct {
	Result EarthquakeResult
	Report_Id string
}

type EarthquakeResult struct {
	Status   string
	Message  string
}

func ParseJsonOf(jsonText string) EarthquakeReport {
	var info EarthquakeReport
	jsonBytes := []byte(jsonText)
	json.Unmarshal(jsonBytes, &info)
	return info
}
