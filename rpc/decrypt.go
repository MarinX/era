package rpc

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Decrypt struct {
	FilePath string   `json:"file_path,omitempty"`
	Text     string   `json:"text,omitempty"`
	Keys     []string `json:"keys"`
	Armor    bool     `json:"armor"`
}

type DecryptResponse struct {
	Output     string `json:"output"`
	OutputPath string `json:"output_path"`
}

func (a *App) DecryptText(req *Decrypt, res *DecryptResponse) error {
	if len(req.Keys) == 0 {
		return errors.New("at least 1 identity needs to be provided")
	}
	if len(req.Text) == 0 {
		return errors.New("text is required")
	}

	var keys []string
	for _, k := range req.Keys {
		var tmp Key
		err := a.store.Read("keys", k, &tmp)
		if err != nil {
			return err
		}
		keys = append(keys, tmp.Private)
	}

	buffer := bytes.NewBuffer(nil)
	err := a.service.Decrypt(strings.NewReader(req.Text), buffer, keys...)
	if err != nil {
		return fmt.Errorf("error decrypting %v", err)
	}
	res.Output = buffer.String()
	return nil
}

func (a *App) DecryptFile(req *Decrypt, res *DecryptResponse) error {
	if len(req.Keys) == 0 {
		return errors.New("at least 1 identity needs to be provided")
	}
	var keys []string
	for _, k := range req.Keys {
		var tmp Key
		err := a.store.Read("keys", k, &tmp)
		if err != nil {
			return err
		}
		keys = append(keys, tmp.Private)
	}
	if _, err := os.Stat(req.FilePath); os.IsNotExist(err) {
		return fmt.Errorf("%s not found", req.FilePath)
	}
	fd, err := os.Open(req.FilePath)
	if err != nil {
		return err
	}
	defer fd.Close()

	res.OutputPath = fd.Name()
	if strings.Contains(res.OutputPath, ".age") {
		res.OutputPath = strings.Replace(res.OutputPath, ".age", "", -1)
	} else {
		res.OutputPath += fmt.Sprintf("-%d", time.Now().Unix())
	}
	out, err := os.Create(res.OutputPath)
	if err != nil {
		return err
	}
	defer out.Close()
	return a.service.Decrypt(fd, out, keys...)
}
