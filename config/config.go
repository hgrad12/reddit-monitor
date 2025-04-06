package config

type Config struct {
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	ClientID   string `yaml:"client_id"`
	Secret     string `yaml:"secret"`
	ServerPort int    `yaml:"server_port"`
	Limit      int    `yaml:"limit"`
}