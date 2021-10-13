package rpc

import (
	"io"
	"net/http"
	"net/rpc"
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

type Config interface {
	GetAPIKey() string
}

type App struct {
	store   Store
	service Service
	config  Config
}

func New(store Store, service Service, config Config) *App {
	return &App{
		store:   store,
		service: service,
		config:  config,
	}
}

func (a *App) ListenAndServe(addr string) error {
	apiKey := a.config.GetAPIKey()
	server := rpc.NewServer()
	if err := server.Register(a); err != nil {
		return err
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-Key")
		if key != apiKey {
			rw.WriteHeader(http.StatusForbidden)
			return
		}

		server.ServeHTTP(rw, r)
	})
	return http.ListenAndServe(addr, mux)
}
