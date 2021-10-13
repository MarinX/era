package service

import (
	"errors"
	"io"
	"strings"

	"filippo.io/age"
	"filippo.io/age/armor"
)

func (s *Service) Encrypt(in io.Reader, out io.Writer, useArmor bool, keys ...string) error {
	var recipients []age.Recipient
	for _, k := range keys {
		key, err := age.ParseRecipients(strings.NewReader(k))
		if err != nil {
			return err
		}
		recipients = append(recipients, key...)
	}
	if len(recipients) == 0 {
		return errors.New("no recipients specified")
	}
	if useArmor {
		a := armor.NewWriter(out)
		defer func() {
			if err := a.Close(); err != nil {
				panic(err)
			}
		}()
		out = a
	}
	w, err := age.Encrypt(out, recipients...)
	if err != nil {
		return err
	}
	if _, err := io.Copy(w, in); err != nil {
		return err
	}
	return w.Close()
}
