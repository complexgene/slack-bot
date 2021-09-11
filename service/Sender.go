package service

import (
	"github.com/slack-go/slack"
	"slack-bot/util"
)

func SendASingleMessage(message string) {
	// Create a new client to slack by giving token
	// Set debug to true while developing
	client := slack.New(util.BotToken, slack.OptionDebug(true))
	// Create the Slack attachment that we will send to the channel
	//attachment := slack.Attachment{
	//	Pretext: "Super Bot Message",
	//	Text:    message,
	//	// Color Styles the Text, making it possible to have like Warnings etc.
	//	Color: "#36a64f",
	//	// Fields are Optional extra data!
	//	//Fields: []slack.AttachmentField{
	//	//	{
	//	//		Title: "Date",
	//	//		Value: time.Now().String(),
	//	//	},
	//	//},
	//}

	// PostMessage will send the message away.
	// First parameter is just the channelID, makes no sense to accept it
	_, _, err := client.PostMessage(
		util.ChannelId,
		// uncomment the item below to add a extra Header to the message, try it out :)
		slack.MsgOptionText(message, false),
		//slack.MsgOptionAttachments(attachment),
	)
	util.HError(err)
}
