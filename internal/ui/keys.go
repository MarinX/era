package ui

import "github.com/MarinX/era/rpc"

func (b *GoBridge) Keys() *ResponseBinder {
	resp := &ResponseBinder{
		Data: new(rpc.KeyResponse),
	}
	keys := new(rpc.KeyResponse)
	err := b.cli.GetKeys(nil, keys)
	if err != nil {
		resp.Error = err.Error()
	} else {
		sortKeysAsc(keys.Keys)
		resp.Data = keys
	}

	return resp
}

func (b *GoBridge) RemoveKey(id string) *ResponseBinder {
	resp := &ResponseBinder{}
	err := b.cli.RemoveKey(&rpc.Key{ID: id}, nil)
	if err != nil {
		resp.Error = err.Error()
	}
	return resp
}

func (b *GoBridge) UpdateKey(id, label string) *ResponseBinder {
	resp := &ResponseBinder{
		Data: new(rpc.KeyResponse),
	}
	err := b.cli.UpdateKey(&rpc.Key{ID: id, Label: label}, resp.Data.(*rpc.KeyResponse))
	if err != nil {
		resp.Error = err.Error()
	}
	return resp
}

func (b *GoBridge) CreateKey(label string) *ResponseBinder {
	resp := &ResponseBinder{
		Data: new(rpc.KeyResponse),
	}
	err := b.cli.GenerateKey(&rpc.Key{Label: label}, resp.Data.(*rpc.KeyResponse))
	if err != nil {
		resp.Error = err.Error()
	}
	return resp
}
