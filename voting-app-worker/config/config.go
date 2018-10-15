package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	LogDir string   `yaml:"log_dir"`
	PgDB   PgConfig `yaml:"pgdb"`
}

type PgConfig struct {
	DBName   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

var AppConfig *Config

func init() {
	AppConfig = newConfig()
}

func newConfig() *Config {
	filename := "./config.yaml"
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Error("Cannot read config file '", filename, "'")
		return nil
	}
	var config = &Config{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		logrus.Error("Cannot parse config file '", filename, "'")
		return nil
	}

	return config
}
