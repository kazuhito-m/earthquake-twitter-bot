package earthquake

import (
	"encoding/json"
)

type EarthquakeReport struct {
	Result        EarthquakeResult
	Report_Id     string
	Region_Name   string
	Calcintensity string
	Report_Time   string
	Request_Time  string
}

type EarthquakeResult struct {
	Status  string
	Message string
}

func ParseJsonOf(jsonText string) EarthquakeReport {
	var info EarthquakeReport
	jsonBytes := []byte(jsonText)
	json.Unmarshal(jsonBytes, &info)
	return info
}

func (this EarthquakeReport) Exists() bool {
	return this.Result.Message != "データがありません"
}

func (this EarthquakeReport) Same(other EarthquakeReport) bool {
	return this.Report_Id == other.Report_Id
}
