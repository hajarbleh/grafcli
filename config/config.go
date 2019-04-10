package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/ghodss/yaml"

	homedir "github.com/mitchellh/go-homedir"
)

// Config variables.
var (
	URL    string
	APIKey string
)

func init() {
	homeDir, err := homedir.Dir()
	if err != nil {
		fmt.Println(errors.Wrap(err, "Cannot get home directory"))
		os.Exit(1)
	}

	defaultConfigPath := homeDir + "/.grafcli/config"
	dat, err := ioutil.ReadFile(defaultConfigPath)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Fatal error config file"))
		os.Exit(1)
	}

	c := map[string]string{}
	err = yaml.Unmarshal(dat, &c)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Unmarshall"))
		os.Exit(1)
	}

	URL = c["url"]
	APIKey = c["api_key"]
}
