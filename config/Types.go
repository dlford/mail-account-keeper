package config

type Config struct {
	AccountConfigs []AccountConfig `yaml:"accounts"`
	Schedule       string          `yaml:"schedule"`
}

type AccountConfig struct {
	Title    string `yaml:"title"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	TLS      bool   `yaml:"tls"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
