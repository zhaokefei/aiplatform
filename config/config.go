package config

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Config struct {
	Redis RedisConfig `yaml:"redis"`
	Mysql MysqlConfig `yaml:"mysql"`
}

const CONFIG_PATH = "config.yaml"


func LoadConfig() (cfg Config, err error) {
	var f *os.File
	path := os.Getenv("GEE_CONFIG_PATH")
	log.Println("CONFIG_PATH", path)
	if path != "" {
		f, err = os.Open(path)
	} else {
		f, err = os.Open(CONFIG_PATH)
	}

	if err != nil {
		return cfg, err
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return cfg, errors.New("unmarshal config file failed")
	}

	log.Println("load config success", cfg)
	return cfg, nil
}

