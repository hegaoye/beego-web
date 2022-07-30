package models

type Bot struct {
	Id          string `json:"id"`
	TgBotId     string `json:"tgBotId"`
	Name        string `json:"name"`
	TgToken     string `json:"tgToken"`
	TgGroupId   string `json:"tgGroupId"`
	GameSummary string `json:"gameSummary"`
	TgGameCode  string `json:"tgGameCode"`
	Domain      string `json:"domain"`
}
