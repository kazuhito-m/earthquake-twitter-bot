package bot

import (
	"github.com/kazuhito-m/earthquake-twitter-bot/config"
	"github.com/kazuhito-m/earthquake-twitter-bot/earthquake"
	"strconv"
)

type MessageMaker struct {
	setting config.BotSetting
}

func (m MessageMaker) Message(report earthquake.EarthquakeReport) string {
	// TODO 実装
	message := "ReportId:" + report.Report_Id
	message += ",ReportTime" + report.Report_Time
	message += ",RequestTime:" + report.Request_Time
	message += ",震度(Calcintensity):" + strconv.Itoa(report.Sindo())
	message += ",緯度(latitude):" + strconv.FormatFloat(report.LatitudeF64(), 'f', 4, 64)
	message += ",軽度(longitude):" + strconv.FormatFloat(report.LongitudeF64(), 'f', 4, 64)
	message += ",地域(RegionName):" + report.Region_Name
	message += ",マグニチュード:" + strconv.FormatFloat(report.MagunitudeF64(), 'f', 4, 64)

	return message
}

func (m MessageMaker) MessageOfNearAndFrequent(nowReport earthquake.EarthquakeReport, lastReport earthquake.EarthquakeReport) string {
	return "同一地域(" + nowReport.Region_Name + ")で地震が頻発している気がします。余震ですかね…気をつけて！"
}

func CreateMessageMaker(setting config.BotSetting) MessageMaker {
	return MessageMaker{setting}
}
