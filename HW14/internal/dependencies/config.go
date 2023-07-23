package dependencies

import (
	"gopkg.in/yaml.v2"
	"os"
)

var Cfg Config

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		DatabaseLocation string `yaml:"databaseLocation"`
	}
}

func LoadConfig(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic("Config file can not be opened")
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Cfg)
}
