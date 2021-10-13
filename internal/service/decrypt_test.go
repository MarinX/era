package service

import (
	"bytes"
	"strings"
	"testing"
)

func TestDecrypt(t *testing.T) {
	keys := map[string]string{
		"age1vl5csek7565t2pugr930n5a9ma26wrer0egc4vsfutg9ehl0pcfqu9x25m": "AGE-SECRET-KEY-1WNXJUVVRTMCKUJU6NQN9ALGCLFX623G0PKWZRDHHL0A0AXUKM5MSJWTW2Y",
		"age1gdkwvvjtzymuuawgd69m7r6ce58wx4gaqtt66rw3xtk5huy22gaqmpw7je": "AGE-SECRET-KEY-1JC6HK6USC67ZZV7G9G3MNGYYETWLP9QP34UGEPRVL52RNZJUUYNSZ2CPPU",
	}

	service := New()
	out := bytes.NewBuffer(nil)

	// lets encrypt a text
	err := service.Encrypt(strings.NewReader("Hello World"), out, true, "age1vl5csek7565t2pugr930n5a9ma26wrer0egc4vsfutg9ehl0pcfqu9x25m")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("Encrypted text %v\n", out.String())
	encTxt := out.String()
	out = bytes.NewBuffer(nil)

	// lets decrypt it
	err = service.Decrypt(strings.NewReader(encTxt), out, keys["age1vl5csek7565t2pugr930n5a9ma26wrer0egc4vsfutg9ehl0pcfqu9x25m"])
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("Decrypted text %v\n", out.String())
	if out.String() != "Hello World" {
		t.Errorf("Expected Hello World, got %v\n", out.String())
		return
	}
}
