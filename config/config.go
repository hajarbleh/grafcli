package config

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const defaultConfigPath = "/home/john/.grafcli/config"

type Config struct {
	Url    string `yaml:"url"`
	ApiKey string `yaml:"api_key"`
}

func Read() (*Config, error) {
	dat, err := ioutil.ReadFile(defaultConfigPath)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Fatal error config file: %s \n", err))
	}
	c := new(Config)
	err = yaml.Unmarshal(dat, &c)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Unmarshall: %s \n", err))
	}
	return c, nil
}
