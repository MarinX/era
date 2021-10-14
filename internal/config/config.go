package config

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/goccy/go-yaml"
)

const (
	ThemeDark  = "dark"
	ThemeLight = "light"
)

type Config struct {
	RPCEnabled bool   `yaml:"rpc_enabled"`
	RPCKey     string `yaml:"rpc_api_key"`
	RPCListen  string `yaml:"rpc_listen"`
	Theme      string `yaml:"theme"`
	DataDir    string `yaml:"-"`
}

func (c *Config) IsValid() []error {
	var ret []error
	if c.RPCEnabled {
		if len(c.RPCKey) == 0 {
			ret = append(ret, errors.New("[invalid config] RPC enabled but key is not set! Please set a API key."))
		}
		if len(c.RPCListen) == 0 {
			ret = append(ret, errors.New("RPC enabled but addr not set. Please set a rpc_listen"))
		}
	}
	if len(c.Theme) == 0 || (c.Theme != ThemeDark && c.Theme != ThemeLight) {
		ret = append(ret, errors.New("invalid theme. Please set it to light or dark"))
	}
	return ret
}

var Defaults = Config{
	RPCEnabled: false,
	RPCKey:     randomKey(),
	RPCListen:  "127.0.0.1:8847",
	Theme:      ThemeDark,
}

func New(home string) (*Config, error) {
	fullDirPath := path.Join(home, ".erapp")
	if _, err := os.Stat(fullDirPath); os.IsNotExist(err) {
		if err := initDefaultConfig(fullDirPath); err != nil {
			return nil, err
		}
		Defaults.DataDir = fullDirPath
		return &Defaults, nil
	}

	// check if we have a config file
	if _, err := os.Stat(path.Join(fullDirPath, "config.yaml")); os.IsNotExist(err) {
		if err := initDefaultConfig(fullDirPath); err != nil {
			return nil, err
		}
		Defaults.DataDir = fullDirPath
		return &Defaults, nil
	}

	cfg := new(Config)
	if err := readConfig(fullDirPath, cfg); err != nil {
		return nil, err
	}
	cfg.DataDir = fullDirPath
	return cfg, nil
}

func initDefaultConfig(dirpath string) error {
	err := os.Mkdir(dirpath, 0755)
	if err != nil {
		return err
	}
	buff, err := yaml.Marshal(Defaults)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path.Join(dirpath, "config.yaml"), buff, 0666)
	if err != nil {
		return err
	}
	return nil
}

func readConfig(dirpath string, cfg *Config) error {
	buff, err := ioutil.ReadFile(path.Join(dirpath, "config.yaml"))
	if err != nil {
		return err
	}
	return yaml.Unmarshal(buff, cfg)
}

func randomKey() string {
	b := make([]byte, 10)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
