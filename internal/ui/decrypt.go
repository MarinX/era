package ui

import "github.com/MarinX/era/rpc"

func (b *GoBridge) DecryptFile(keys []interface{}, data string) *ResponseBinder {
	resp := &ResponseBinder{
		Data: new(rpc.DecryptResponse),
	}
	var rep []string
	for _, c := range keys {
		rep = append(rep, c.(string))
	}

	err := b.cli.DecryptFile(&rpc.Decrypt{FilePath: data, Keys: rep}, resp.Data.(*rpc.DecryptResponse))
	if err != nil {
		resp.Error = err.Error()
	}
	return resp
}

func (b *GoBridge) DecryptText(keys []interface{}, data string) *ResponseBinder {
	resp := &ResponseBinder{
		Data: new(rpc.DecryptResponse),
	}
	var rep []string
	for _, c := range keys {
		rep = append(rep, c.(string))
	}
	err := b.cli.DecryptText(&rpc.Decrypt{Text: data, Keys: rep}, resp.Data.(*rpc.DecryptResponse))
	if err != nil {
		resp.Error = err.Error()
	}
	return resp
}
