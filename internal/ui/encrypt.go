package ui

import "github.com/MarinX/era/rpc"

func (b *GoBridge) EncryptFile(contacts []interface{}, data string, armor bool) *ResponseBinder {
	resp := &ResponseBinder{
		Data: new(rpc.EncryptResponse),
	}
	var rep []string
	for _, c := range contacts {
		rep = append(rep, c.(string))
	}

	err := b.cli.EncryptFile(&rpc.Encrypt{FilePath: data, Recipients: rep, Armor: armor}, resp.Data.(*rpc.EncryptResponse))
	if err != nil {
		resp.Error = err.Error()
	}
	return resp
}

func (b *GoBridge) EncryptText(contacts []interface{}, data string, armor bool) *ResponseBinder {
	resp := &ResponseBinder{
		Data: new(rpc.EncryptResponse),
	}
	var rep []string
	for _, c := range contacts {
		rep = append(rep, c.(string))
	}
	err := b.cli.EncryptText(&rpc.Encrypt{Text: data, Recipients: rep, Armor: armor}, resp.Data.(*rpc.EncryptResponse))
	if err != nil {
		resp.Error = err.Error()
	}
	return resp
}
