package config

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/goccy/go-yaml"
)

func TestConfig(t *testing.T) {
	defer os.RemoveAll(".erapp")

	cfg, err := New("")
	if err != nil {
		t.Error(err)
		return
	}

	errors := cfg.IsValid()
	for _, e := range errors {
		t.Error(e)
	}

	fullDirPath := path.Join("", ".erapp")
	if _, err := os.Stat(fullDirPath); os.IsNotExist(err) {
		t.Error("config should be set but got empty folder")
		return
	}

	// modify
	cfg.RPCKey = "Hello"
	buff, err := yaml.Marshal(cfg)
	if err != nil {
		t.Error(err)
		return
	}
	ioutil.WriteFile(path.Join(fullDirPath, "config.yaml"), buff, 0666)

	// read again
	cfg, err = New("")
	if err != nil {
		t.Error(err)
		return
	}
	if cfg.RPCKey != "Hello" {
		t.Error("rpckey did not update on second load")
		return
	}

	cfg.RPCEnabled = true
	cfg.RPCKey = ""
	errors = cfg.IsValid()
	if len(errors) == 0 {
		t.Error("should show error for invalid rpc key but got none")
	}

	cfg.RPCEnabled = false
	cfg.Theme = "bla"
	errors = cfg.IsValid()
	if len(errors) == 0 {
		t.Error("should show error for invalid theme but got none")
	}
}
