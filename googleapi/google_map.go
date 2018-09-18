package googleapi

import (
	"fmt"
	"github.com/kazuhito-m/earthquake-twitter-bot/web"
)

const MAP_API_URL_TEMPLATE = "https://maps.googleapis.com/maps/api/geocode/json?language=ja&region=jp&sensor=false&key=%v&latlng=%v,%v"

type GoogleMap struct {
	client web.Client
}

func (m GoogleMap) GeoCode(latitude float64, longitude float64) GeoCode {
	url := CreateUrl(latitude, longitude)
	jsonText := m.client.Get(url)
	return ParseJsonOf(jsonText)
}

func CreateUrl(latitude float64, longitude float64) string {
	return fmt.Sprintf(MAP_API_URL_TEMPLATE, "test", latitude, longitude)
}

func CreateGoogleMap(client web.Client) GoogleMap {
	return GoogleMap{client}
}
