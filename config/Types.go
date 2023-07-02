package config

type Config struct {
	AccountConfigs []AccountConfig
}

type AccountConfig struct {
	Title    string `json:"title"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	TLS      bool   `json:"tls"`
	Username string `json:"username"`
	Password string `json:"password"`
	MailTo   string `json:"mailTo"`
	Schedule string `json:"schedule"`
}
