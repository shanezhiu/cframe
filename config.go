package main

import (
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

type Config struct {
	Local *Node   `toml:"local"`
	Peers []*Node `toml:"peers"`
}

type Node struct {
	Addr string `toml:"addr"`
	CIDR string `toml:"cidr"`
}

func ParseConfig(path string) (*Config, error) {
	cnt, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = toml.Unmarshal(cnt, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
