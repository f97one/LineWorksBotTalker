package settings

type LWBotTalkConfig struct {
	ApiId          string `json:"api_id"`
	ConsumerKey    string `json:"consumer_key"`
	ServerId       string `json:"server_id"`
	PrivateKeyPath string `json:"private_key_path"`
	BotNo          int    `json:"bot_no"`
}

var config *LWBotTalkConfig

func GetConfig() LWBotTalkConfig {
	return *config
}

func SetConfig(conf *LWBotTalkConfig) {
	config = conf
}
