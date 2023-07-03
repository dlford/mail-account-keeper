package config

type Config struct {
	AccountConfigs []AccountConfig
	AlertConfig    AlertConfig
}

type AccountConfig struct {
	Title    string `json:"title"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	MailTo   string `json:"mail_to"`
	Schedule string `json:"schedule"`
}

type AlertConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	MailTo   string `json:"mail_to"`
}
