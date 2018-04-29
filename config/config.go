package config

type CiscoConfig struct {
	Server   string `ini:"server"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type CbsConfig struct {
	Url string `ini:"url"`
}

type Config struct {
	Cisco CiscoConfig `ini:"cisco"`
	CBS   CbsConfig   `ini:"cbs"`
}

const CONFIG_FILE = "package_config.ini"

/* --- Keep track of current config globally --- */
var currentConfig Config

func GetConfig() *Config {
	return &currentConfig
}
