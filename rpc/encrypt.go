package rpc

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Encrypt struct {
	FilePath   string   `json:"file_path,omitempty"`
	Text       string   `json:"text,omitempty"`
	Recipients []string `json:"recipients"`
	Armor      bool     `json:"armor"`
}

type EncryptResponse struct {
	OutputPath string `json:"output_path,omitempty"`
	Output     string `json:"output,omitempty"`
}

func (a *App) EncryptText(req *Encrypt, res *EncryptResponse) error {
	if len(req.Recipients) == 0 {
		return errors.New("at least 1 recipient needs to be added")
	}
	if len(req.Text) == 0 {
		return errors.New("empty body is not allowed")
	}
	var keys []string
	for _, rec := range req.Recipients {
		var tmp Key
		err := a.store.Read("contacts", rec, &tmp)
		if err != nil {
			return err
		}
		keys = append(keys, tmp.Key)
	}
	buff := bytes.NewBuffer(nil)
	err := a.service.Encrypt(strings.NewReader(req.Text), buff, true, keys...)
	if err != nil {
		return err
	}
	res.Output = buff.String()
	return nil
}

func (a *App) EncryptFile(req *Encrypt, res *EncryptResponse) error {
	if len(req.Recipients) == 0 {
		return errors.New("at least 1 recipient needs to be added")
	}
	var keys []string
	for _, rec := range req.Recipients {
		var tmp Key
		err := a.store.Read("contacts", rec, &tmp)
		if err != nil {
			return err
		}
		keys = append(keys, tmp.Key)
	}
	if _, err := os.Stat(req.FilePath); os.IsNotExist(err) {
		return fmt.Errorf("%s not found", req.FilePath)
	}
	fd, err := os.Open(req.FilePath)
	if err != nil {
		return err
	}
	defer fd.Close()
	res.OutputPath = fmt.Sprintf("%s.age", fd.Name())
	out, err := os.Create(res.OutputPath)
	if err != nil {
		return err
	}
	defer out.Close()
	return a.service.Encrypt(fd, out, req.Armor, keys...)
}
