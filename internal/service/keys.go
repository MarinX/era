package service

import (
	"strings"

	"filippo.io/age"
)

func (s *Service) CreateKey() (string, string, error) {
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		return "", "", err
	}
	return identity.String(), identity.Recipient().String(), nil
}

func (s *Service) IsIdentity(key string) bool {
	_, err := age.ParseIdentities(strings.NewReader(key))
	return err == nil
}

func (s *Service) IsRecipient(key string) bool {
	_, err := age.ParseRecipients(strings.NewReader(key))
	return err == nil
}
