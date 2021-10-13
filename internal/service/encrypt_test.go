package service

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestEncrypt(t *testing.T) {
	keys := []string{
		"age1vl5csek7565t2pugr930n5a9ma26wrer0egc4vsfutg9ehl0pcfqu9x25m",
		"age1gdkwvvjtzymuuawgd69m7r6ce58wx4gaqtt66rw3xtk5huy22gaqmpw7je",
		"age1x78chxt800rx9y2wuuzzc8qfzqy9nljreq92scm2rj5a43xtqfsqdr85xu",
	}

	service := New()

	t.Log("Should encrypt using single key")
	buff := bytes.NewBuffer(nil)
	err := service.Encrypt(strings.NewReader("Hello World"), buff, true, keys[0])
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Should encrypt using multiple key")
	buff = bytes.NewBuffer(nil)
	err = service.Encrypt(strings.NewReader("Hello World"), buff, true, keys...)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Should fail on using invalid key")
	buff = bytes.NewBuffer(nil)
	err = service.Encrypt(strings.NewReader("Hello World"), buff, true, "bla")
	if err == nil {
		t.Error("Should fail encrypting but got nil error")
		return
	}

	t.Log("Should encrypt to a file")
	fd, err := os.CreateTemp(os.TempDir(), "encrypted.txt")
	if err != nil {
		t.Error(err)
		return
	}
	err = service.Encrypt(strings.NewReader("Hello World"), fd, true, keys[0])
	if err != nil {
		t.Error(err)
		return
	}
	fd.Close()
	defer os.Remove(fd.Name())
	fileOutput, err := ioutil.ReadFile(fd.Name())
	if err != nil {
		t.Error(err)
		return
	}
	if !strings.Contains(string(fileOutput), "-----BEGIN AGE ENCRYPTED FILE-----") {
		t.Error("invalid file output")
	}
}
