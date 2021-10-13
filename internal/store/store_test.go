package store

import (
	"os"
	"testing"
)

func TestStore(t *testing.T) {
	str, err := New(os.TempDir())
	if err != nil {
		t.Error(err)
		return
	}
	defer str.Close()
	defer os.Remove(str.location)

	t.Logf("DB Path %v\n", str.location)

	type fakeModel struct {
		A string   `json:"a"`
		B int      `json:"b"`
		C []string `json:"c"`
	}
	model := fakeModel{
		A: "aaa",
		B: 2,
		C: []string{"one", "two"},
	}

	err = str.Create("contacts", "123", model)
	if err != nil {
		t.Error(err)
		return
	}

	err = str.Create("keys", "123", model)
	if err != nil {
		t.Error(err)
		return
	}

	var ff fakeModel
	// not found
	err = str.Read("contacts", "bla", &ff)
	if err == nil || err.Error() != "bla not found" {
		t.Errorf("expecte error but got %v\n", err)
	}

	// read back the data
	err = str.Read("contacts", "123", &ff)
	if err != nil {
		t.Error(err)
		return
	}
	if ff.B != 2 {
		t.Errorf("Expeting 2 from model, but got %v\n", ff.B)
		return
	}

	// lets update
	ff.B = 100
	err = str.Update("contacts", "123", &ff)
	if err != nil {
		t.Error(err)
		return
	}

	ff = fakeModel{}
	err = str.Read("contacts", "123", &ff)
	if err != nil {
		t.Error(err)
		return
	}
	if ff.B != 100 {
		t.Errorf("expect 100 after update but got %v", ff.B)
	}

	// remove record
	err = str.Delete("contacts", "123")
	if err != nil {
		t.Error(err)
		return
	}
	// not found
	err = str.Read("contacts", "123", &ff)
	if err == nil || err.Error() != "123 not found" {
		t.Errorf("expecte error but got %v\n", err)
	}

	// should not panic even if we dont have that key
	err = str.Delete("contacts", "24")
	if err != nil {
		t.Error(err)
		return
	}

}
