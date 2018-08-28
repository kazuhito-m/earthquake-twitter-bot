package earthquake

import (
	"github.com/kazuhito-m/earthquake-twitter-bot/web"
	"strings"
	"time"
)

const SITE_URL_TEMPLATE = "http://www.kmoni.bosai.go.jp/new/webservice/hypo/eew/{time_id}.json"

type Site struct {
	client web.Client
}

func CreateTimeId() string {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	return time.Now().In(loc).Format("20060102150405")
}

func createUrl() string {
	return strings.Replace(SITE_URL_TEMPLATE, "{time_id}", CreateTimeId(), 1)
}

func (site Site) GetNowEarthquakeAnalyze() EarthquakeInformation {
	url := createUrl()
	jsonText := site.client.Get(url)
	return ParseJsonOf(jsonText)
}

func CreateSite(client web.Client) Site {
	return Site{client}
}
