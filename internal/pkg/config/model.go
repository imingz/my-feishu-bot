package config

type Config struct {
	APP         AppConfig         `yaml:"App"`
	Qrcode      QrcodeConfig      `yaml:"qrcode"`
	RoomBalance RoomBalanceConfig `yaml:"roomBalance"`
}

type AppConfig struct {
	ID                string `yaml:"ID"`
	Secret            string `yaml:"Secret"`
	VerificationToken string `yaml:"VerificationToken"`
	EventEncryptKey   string `yaml:"EventEncryptKey"`
}

type QrcodeConfig struct {
	OpenId  string `yaml:"openId"`  // 慧湖通的 openId
	BaseUrl string `yaml:"baseUrl"` // 慧湖通的 BaseUrl
	ChatId  string `yaml:"chatId"`  // 监听的群聊 chatId
}

type RoomBalanceConfig struct {
	OpenId string `yaml:"openId"` // 接受消息的人的 openId
}
