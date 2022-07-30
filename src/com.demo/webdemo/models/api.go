package models

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
	"strings"
)

const (
	// GetBotUrl 獲取機器人 get
	GetBotUrl = "{host}/bot/load/name/{botName}"
	// GetGroupByIdUrl 根據群id後去群信息
	GetGroupByIdUrl = "{host}/group/load/id/{groupId}"
)

var bot *Bot
var ServerHost string

func GetBot() Bot {
	return *bot
}

// 機器人url
func botUrl(serverHost string, botName string) string {
	url := strings.Replace(GetBotUrl, "{host}", serverHost, -1)
	url = strings.Replace(url, "{botName}", botName, -1)
	return url
}

// 群id url
func groupByIdUrl(serverHost string, groupId string) string {
	url := strings.Replace(GetGroupByIdUrl, "{host}", serverHost, -1)
	url = strings.Replace(url, "{groupId}", groupId, -1)
	return url
}

func initData(serverHost string, botName string) {
	//獲取bot信息
	botUrl := botUrl(serverHost, botName)
	req := httplib.Get(botUrl)
	result, err := req.String()
	if err != nil {
	}

	var r R
	json.Unmarshal([]byte(result), &r)
	j, _ := json.Marshal(r.Data)
	json.Unmarshal(j, &bot)

	//獲取群組信息
	id := bot.Id
	groupUrl := groupByIdUrl(serverHost, id)
	reqGroup := httplib.Get(groupUrl)
	resultOfGroup, _ := reqGroup.String()
	json.Unmarshal([]byte(resultOfGroup), &r)
	jGroup, _ := json.Marshal(r.Data)
	json.Unmarshal(jGroup, &bot)
	logs.Debug("機器人完整信息 %s", bot)
}

//初始化讀取配置
func init() {

	serverHost := beego.AppConfig.String("sys::server_host")
	botName := beego.AppConfig.String("sys::bot_name")

	ServerHost = serverHost

	logs.Debug("服務器地址 %s", serverHost)
	logs.Debug("機器人名稱 %s", botName)

	initData(serverHost, botName)
}
