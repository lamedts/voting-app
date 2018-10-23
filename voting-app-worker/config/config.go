package config

import (
	"io/ioutil"
	"os"

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
	if env := os.Getenv("PG_DBNAME"); env != "" {
		config.PgDB.DBName = os.Getenv("PG_DBNAME")
	}
	if env := os.Getenv("PG_USER"); env != "" {
		config.PgDB.User = os.Getenv("PG_USER")
	}
	if env := os.Getenv("PG_PW"); env != "" {
		config.PgDB.Password = os.Getenv("PG_PW")
	}
	if env := os.Getenv("PG_HOST"); env != "" {
		config.PgDB.Host = os.Getenv("PG_HOST")
	}
	if env := os.Getenv("PG_PORT"); env != "" {
		config.PgDB.Port = os.Getenv("PG_PORT")
	}
	return config
}
