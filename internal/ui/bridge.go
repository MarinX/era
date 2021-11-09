package ui

import (
	"github.com/MarinX/era/internal/config"
	"github.com/MarinX/era/rpc"
	"github.com/atotto/clipboard"
	"github.com/wailsapp/wails"
)

type GoBridge struct {
	runtime *wails.Runtime
	cfg     *config.Config
	cli     *rpc.App
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

func (b *GoBridge) WailsInit(runtime *wails.Runtime) error {
	b.runtime = runtime
	return nil
}

func (b *GoBridge) CopyToClipboard(text string) *ResponseBinder {
	resp := &ResponseBinder{}
	err := clipboard.WriteAll(text)
	if err != nil {
		resp.Error = err.Error()
	}
	return resp
}

func (b *GoBridge) Settings() *ResponseBinder {
	return &ResponseBinder{
		Data: b.cfg,
	}
}

func (b *GoBridge) SelectFile() *ResponseBinder {
	file := b.runtime.Dialog.SelectFile("Select file")
	return &ResponseBinder{
		Data: file,
	}
}
