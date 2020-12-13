package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

const filePath = "file/configurations.json"

type (
	//Config is all configurations in the service
	Config struct {
		Port   PortConfig   `json:"port"`
		Worker WorkerConfig `json:"worker"`
	}

	//PortConfig to get all port configs
	PortConfig struct {
		Main int `json:"main"`
	}

	//WorkerConfig to get concurrent worker configs
	WorkerConfig struct {
		Default int `json:"default"`
	}
)

//Init to initialize configs
func Init() (cfg Config, err error) {
	configs, err := ioutil.ReadFile(filePath)
	if err != nil {
		err = errors.Wrapf(err, "[Config] fail to read file from %s", filePath)
		return
	}

	err = json.Unmarshal(configs, &cfg)
	if err != nil {
		err = errors.Wrapf(err, "[Config] fail to unmarshal file from %s", filePath)
		return
	}

	return
}
