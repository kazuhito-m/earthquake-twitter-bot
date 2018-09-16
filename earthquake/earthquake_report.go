package earthquake

import (
	"encoding/json"
)

type EarthquakeReport struct {
	Result        EarthquakeResult
	Report_Id     string
	Calcintensity json.Number
	Region_Name   string
	Report_Time   string
	Latitude      json.Number
	Longitude     json.Number
	Magunitude    json.Number
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

// Utility functions

func convJsonNumberToFloat(number json.Number) float64 {
	value, err := number.Float64()
	if err != nil {
		return 0
	}
	return value
}

// Custom fields

func (this EarthquakeReport) Sindo() int {
	value, err := this.Calcintensity.Int64()
	if err != nil {
		return 0
	}
	return int(value)
}

func (this EarthquakeReport) LatitudeF64() float64 {
	return convJsonNumberToFloat(this.Latitude)
}

func (this EarthquakeReport) LongitudeF64() float64 {
	return convJsonNumberToFloat(this.Longitude)
}

func (this EarthquakeReport) MagunitudeF64() float64 {
	return convJsonNumberToFloat(this.Magunitude)
}
