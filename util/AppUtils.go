package util

import (
	"context"
	"github.com/labstack/gommon/log"
	"slack-bot/models"
	"slack-bot/repo"
)

// BotToken Init Variables required Later
var BotToken = "BOT_TOKEN"
var ChannelId = "CHANNEL_ID"

func HError(e error) {
	if e != nil {
		log.Error("Got an error: " + e.Error())
		return
	}
}

func InitBotValues() {
	db := repo.GetDBConn()
	utilKVs := make(map[string]string)
	botDetailsList := make([]models.BotDetails, 0)
	err := db.NewSelect().Model(&botDetailsList).Scan(context.Background())
	HError(err)
	for _, botDetail := range botDetailsList {
		utilKVs[botDetail.Key] = botDetail.Value
	}
	// Populate Init Variables required Later
	BotToken = utilKVs[BotToken]
	ChannelId = utilKVs[ChannelId]
}
