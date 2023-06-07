package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	FullNode FullNode `toml:"fullnode"`
	Pruner   Pruner   `toml:"pruner"`
	Aws      Aws      `toml:"aws"`
}

type FullNode struct {
	Path       string
	Data_Path  string
	Chain_Name string
}

type Pruner struct {
	Path string
}

type Aws struct {
	Region string
	Bucket string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load(configFile string) {
	if _, err := toml.DecodeFile(configFile, c); err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}
