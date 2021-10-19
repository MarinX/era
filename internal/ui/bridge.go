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
