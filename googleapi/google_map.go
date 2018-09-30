package googleapi

import (
	"fmt"
	"github.com/kazuhito-m/earthquake-twitter-bot/config"
	"github.com/kazuhito-m/earthquake-twitter-bot/web"
)

const MAP_API_URL_TEMPLATE = "https://maps.googleapis.com/maps/api/geocode/json?language=ja&region=jp&sensor=false&key=%v&latlng=%v,%v"

type GoogleMap struct {
	client   web.Client
	settings config.GoogleApiSettings
}

func (m GoogleMap) GeoCode(latitude float64, longitude float64) GeoCode {
	url := CreateUrl(latitude, longitude, m.settings)
	jsonText := m.client.Get(url)
	return ParseJsonOf(jsonText)
}

func CreateUrl(latitude float64, longitude float64, settings config.GoogleApiSettings) string {
	return fmt.Sprintf(MAP_API_URL_TEMPLATE, settings.ApiKey, latitude, longitude)
}

func CreateGoogleMap(client web.Client, settings config.GoogleApiSettings) GoogleMap {
	return GoogleMap{client, settings}
}
