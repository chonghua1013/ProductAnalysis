package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
)

type Config struct {
	ServerAddress string `yaml:"server_address"`
	DB            DatabaseConfig
	JWT           JWTConfig
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type JWTConfig struct {
	Secret string `yaml:"secret"`
	Expire int    `yaml:"expire"` // in hours
}

func LoadConfig() (*Config, error) {
	// 读取配置文件
	configFile, err := os.ReadFile("configs/app.yaml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(configFile, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c DatabaseConfig) GetDSN() string {
	return "host=" + c.Host + " user=" + c.User + " password=" + c.Password +
		" dbname=" + c.Name + " port=" + strconv.Itoa(c.Port) + " sslmode=disable TimeZone=Asia/Shanghai"
}
