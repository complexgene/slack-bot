package models

type Event struct {
	Id               int    `json:"id"`
	EventTitle       string `json:"event_title"`
	EventDescription string `json:"event_description"`
}

type BotDetails struct {
	Id    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SlackReq struct {
	Challenge string `json:"challenge"`
}

type BotEvent struct {
	Token    string `json:"token"`
	TeamId   string `json:"team_id"`
	ApiAppId string `json:"api_app_id"`
	Event    struct {
		Type    string `json:"type"`
		User    string `json:"user"`
		Text    string `json:"text"`
		Ts      string `json:"ts"`
		Channel string `json:"channel"`
		EventTs string `json:"event_ts"`
	} `json:"event"`
	Type        string   `json:"type"`
	EventId     string   `json:"event_id"`
	EventTime   int64    `json:"event_time"`
	AuthedUsers []string `json:"authed_users"`
}
