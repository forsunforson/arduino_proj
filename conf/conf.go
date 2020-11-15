package conf

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

// GlobalConf config from yaml
type GlobalConf struct {
	Server struct {
		Mqtt struct {
			broker       string `yaml:"broker"`
			client       string `yaml:"client"`
			username     string `yaml:"username"`
			password     string `yaml:"password"`
			cleanSession bool   `yaml:"cleansession"`
		}
	}
}

var (
	globalConf *GlobalConf
)

// GetGlobalConf read conf from yaml file
func GetGlobalConf() (*GlobalConf, error) {
	if globalConf == nil {
		b, err := ioutil.ReadFile("./conf.yaml")
		if err != nil {
			return nil, err
		}
		conf := &GlobalConf{}
		err = yaml.Unmarshal(b, conf)
		if err != nil {
			return nil, err
		}
		globalConf = conf
	}

	return globalConf, nil
}
