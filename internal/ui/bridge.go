package ui

import (
	"github.com/MarinX/era/internal/config"
	"github.com/MarinX/era/rpc"
	"github.com/atotto/clipboard"
)

type GoBridge struct {
	cfg *config.Config
	cli *rpc.App
}

type ResponseBinder struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func New(cfg *config.Config, cli *rpc.App) *GoBridge {
	return &GoBridge{
		cfg: cfg,
		cli: cli,
	}
}

func (b *GoBridge) Contacts() *ResponseBinder {
	resp := &ResponseBinder{
		Data: new(rpc.ContactResponse),
	}

	cnts := new(rpc.ContactResponse)
	err := b.cli.GetContacts(nil, cnts)
	if err != nil {
		resp.Error = err.Error()
	} else {
		sortContactsAsc(cnts.Contacts)
		resp.Data = cnts
	}

	return resp
}

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

func (b *GoBridge) CopyToClipboard(text string) *ResponseBinder {
	resp := &ResponseBinder{}
	err := clipboard.WriteAll(text)
	if err != nil {
		resp.Error = err.Error()
	}
	return resp
}

func (b *GoBridge) APIKey() string {
	return b.cfg.RPCKey
}

func (b *GoBridge) RPCEnabled() bool {
	return b.cfg.RPCEnabled
}
