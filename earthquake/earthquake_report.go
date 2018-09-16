package earthquake

import (
	"encoding/json"
	"math"
	"time"
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

func (this EarthquakeReport) NearAndFrequent(other EarthquakeReport, frequentThresholdHour int) bool {
	if this.Region_Name != other.Region_Name {
		return false
	}
	myTime := this.reportTime()
	otherTime := other.reportTime()
	timeNothing := timeNothing()
	if myTime.Equal(timeNothing) || otherTime.Equal(timeNothing) {
		return false
	}
	intervalSeconds := math.Abs(myTime.Sub(otherTime).Seconds())
	frequentThresholdSeconds := toSec(frequentThresholdHour)
	return intervalSeconds > frequentThresholdSeconds
}

// Utility functions

func convertJsonNumberToFloat(number json.Number) float64 {
	value, err := number.Float64()
	if err != nil {
		return 0
	}
	return value
}

func timeNothing() time.Time {
	return time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
}

func toSec(hours int) float64 {
	return float64(time.Duration(hours) * time.Minute * time.Hour)
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
	return convertJsonNumberToFloat(this.Latitude)
}

func (this EarthquakeReport) LongitudeF64() float64 {
	return convertJsonNumberToFloat(this.Longitude)
}

func (this EarthquakeReport) MagunitudeF64() float64 {
	return convertJsonNumberToFloat(this.Magunitude)
}

func (this EarthquakeReport) reportTime() time.Time {
	value, err := time.Parse(this.Request_Time, "2006/01/02 15:04:05")
	if err == nil {
		return timeNothing()
	}
	return value
}
