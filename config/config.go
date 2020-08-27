package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// ConfigFile is the default config file
var ConfigFile = "./config.yml"

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
	Addr    string
	Mode    string
	Version string
	Secret  string
}

type DatabaseConfig struct {
	Host    string
	Port    int32
	Timeout int32
}

// global configs
var (
	Global   Config
	Server   ServerConfig
	Database DatabaseConfig
)

// Load config from file
func Load(file string) (Config, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("%v", err)
		return Global, err
	}

	err = yaml.Unmarshal(data, &Global)
	if err != nil {
		log.Printf("%v", err)
		return Global, err
	}

	Server = Global.Server
	Database = Global.Database

	return Global, nil
}

// loads configs
func init() {
	if os.Getenv("config") != "" {
		ConfigFile = os.Getenv("config")
	}
	Load(ConfigFile)
}
