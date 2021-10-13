package service

import (
	"bufio"
	"io"
	"strings"

	"filippo.io/age"
	"filippo.io/age/armor"
)

func (s *Service) Decrypt(in io.Reader, out io.Writer, keys ...string) error {
	rr := bufio.NewReader(in)
	if start, _ := rr.Peek(len(armor.Header)); string(start) == armor.Header {
		in = armor.NewReader(rr)
	} else {
		in = rr
	}
	var identities []age.Identity
	for _, name := range keys {
		ids, err := age.ParseIdentities(strings.NewReader(name))
		if err != nil {
			return err
		}
		identities = append(identities, ids...)
	}
	r, err := age.Decrypt(in, identities...)
	if err != nil {
		return err
	}
	_, err = io.Copy(out, r)
	return err
}
