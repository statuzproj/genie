package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Target struct {
	Type     string `yaml:"type"`
	Endpoint string `yaml:"endpoint"`
	Interval string `yaml:"interval"`
	Method   string `yaml:"method"`
}

type Endpoint struct {
	Name   string `yaml:"name"`
	Target Target `yaml:"target"`
}

var cfgPath = "/etc/config/"

func LoadEndpoints() (ednpoints []Endpoint, err error) {

	var cfgFile = cfgPath + "endpoints.yaml"
	var endpoints []Endpoint

	content, err := os.ReadFile(cfgFile)
	if err != nil {
		log.Println(err)
		return endpoints, err
	}

	err = yaml.Unmarshal(content, &endpoints)
	if err != nil {
		log.Println(err)
		return endpoints, err
	}
	return endpoints, nil
}
