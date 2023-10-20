package config

import (
	"errors"
	"os"
	"log"

	"gopkg.in/yaml.v3"
)


type RedisConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Password string `yaml:"password"`
	DB int `yaml:"db"`
}

type MysqlConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}


type Config struct {
	Redis RedisConfig `yaml:"redis"`
	Mysql MysqlConfig `yaml:"mysql"`
}

var Cfg Config
const CONFIG_PATH = "config.yaml"

func init() {
	err := Load()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}


func ReLoad() error {
	f, err := os.Open(CONFIG_PATH)
	if err != nil {
		return err
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&Cfg)
	if err != nil {
		return errors.New("unmarshal config file failed")
	}

	log.Println("load config success", Cfg)
	return nil
}


func Load() error {
	if Cfg != (Config{}) {
		return nil
	}
	return ReLoad()
}
