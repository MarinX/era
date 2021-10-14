package rpc

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Store interface {
	Create(key, id string, value interface{}) error
	ReadAll(key string, value interface{}) error
	Read(key, id string, value interface{}) error
	Update(key, id string, value interface{}) error
	Delete(key, id string) error
}

type Service interface {
	Encrypt(in io.Reader, out io.Writer, useArmor bool, keys ...string) error
	Decrypt(in io.Reader, out io.Writer, keys ...string) error
	IsIdentity(key string) bool
	IsRecipient(key string) bool
	CreateKey() (private string, public string, err error)
}

type App struct {
	store   Store
	service Service
}

func New(store Store, service Service) *App {
	return &App{
		store:   store,
		service: service,
	}
}

func ListenAndServe(methods interface{}, addr string, apiKey string) error {
	server := rpc.NewServer()
	if err := server.Register(methods); err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-Key")
		if key != apiKey {
			rw.WriteHeader(http.StatusForbidden)
			return
		}

		rw.Header().Set("Content-type", "application/json")
		server.ServeCodec(jsonrpc.NewServerCodec(&Conn{
			in:  r.Body,
			out: rw,
		}))

	})
	return http.ListenAndServe(addr, mux)
}
