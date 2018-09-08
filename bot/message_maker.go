package bot

import (
	"github.com/kazuhito-m/earthquake-twitter-bot/config"
	"github.com/kazuhito-m/earthquake-twitter-bot/earthquake"
)

type MessageMaker struct {
	setting config.BotSetting
}

func (m MessageMaker) Message(report earthquake.EarthquakeReport) string {
	// TODO 実装
	message := "ReportId:" + report.Report_Id  + ",ReportTime" + report.Report_Time + ",RequestTime:"  + report.Request_Time
	return message
}


func CreateMessageMaker(setting config.BotSetting) MessageMaker {
	return MessageMaker{setting}
}
