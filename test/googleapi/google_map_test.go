package googleapi

import (
	"github.com/kazuhito-m/earthquake-twitter-bot/config"
	"github.com/kazuhito-m/earthquake-twitter-bot/googleapi"
	"strings"
	"testing"
)

func TestパラメータからAPIのURLを組み立てることが出来る(t *testing.T) {
	settings := config.GoogleApiSettings{"AbCxYz"}

	actual := googleapi.CreateUrl(123.456, 987.654, settings)

	if strings.Index(actual, "https://maps.googleapis.com") != 0 {
		t.Errorf("URLの先頭が「GoogleMapAPI」のものではない。値:'%v'", actual)
	}
	if !strings.Contains(actual, "&key=AbCxYz&latlng=123.456,987.654") {
		t.Errorf("URLにキーと緯度経度が含まれていない。値:'%v'", actual)
	}
}
