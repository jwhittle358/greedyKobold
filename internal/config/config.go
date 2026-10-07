package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Name        string `yaml:"name"`
		Environment string `yaml:"environment"`
		LogLevel    string `yaml:"log_level"`
	} `yaml:"app"`

	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`

	Syslog struct {
		UDPPort int `yaml:"udp_port"`
		TCPPort int `yaml:"tcp_port"`
	} `yaml:"syslog"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
