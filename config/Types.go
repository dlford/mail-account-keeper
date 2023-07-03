package config

type Config struct {
	AccountConfigs []AccountConfig
}

type AccountConfig struct {
	Title    string `json:"title"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Email    string `json:"email"`
	Password string `json:"password"`
	MailTo   string `json:"mail_to"`
	Schedule string `json:"schedule"`
}
