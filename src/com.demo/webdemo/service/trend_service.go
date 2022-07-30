package service

import (
	"com.demo/webdemo/models"
	"encoding/json"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/httplib"
	"strings"
)

type TrendService struct {
}

func (receiver TrendService) GetCount() any {
	trendTableUrl := "{host}/trendRecord/count/{group_id}"
	trendTableUrl = strings.Replace(trendTableUrl, "{host}", models.ServerHost, -1)
	trendTableUrl = strings.Replace(trendTableUrl, "{group_id}", models.GetBot().TgGroupId, -1)
	req := httplib.Get(trendTableUrl)
	result, err := req.String()
	if err != nil {
	}
	var r models.R
	json.Unmarshal([]byte(result), &r)

	logs.Debug(r.Data)

	return r.Data
}
