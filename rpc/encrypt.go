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
	// check if we have good recipients
	for _, rec := range req.Recipients {
		if !a.service.IsRecipient(rec) {
			return fmt.Errorf("%s is not a valid recipient key", rec)
		}
	}
	buff := bytes.NewBuffer(nil)
	err := a.service.Encrypt(strings.NewReader(req.Text), buff, true, req.Recipients...)
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
	// check if we have good recipients
	for _, rec := range req.Recipients {
		if !a.service.IsRecipient(rec) {
			return fmt.Errorf("%s is not a valid recipient key", rec)
		}
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
	return a.service.Encrypt(fd, out, req.Armor, req.Recipients...)
}
