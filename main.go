package main

import (
	"fmt"
	"slack-bot/http"
	"slack-bot/util"
)

func main() {
	fmt.Println("Starting Slack Bot..")
	util.InitBotValues()
	http.InitRoutes()
}
