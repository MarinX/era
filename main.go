package main

import (
	_ "embed"
	"flag"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MarinX/era/internal/config"
	"github.com/MarinX/era/internal/service"
	"github.com/MarinX/era/internal/ui"
	"github.com/MarinX/era/rpc"
	"github.com/MarinX/era/store"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails"
)

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

//go:embed frontend/dist/index.html
var html string

var noui = flag.Bool("noui", false, "Start Era without UI (probably you want RPC enabled)")
var debug = flag.Bool("debug", false, "enable debug and inspector mode")

func main() {
	rand.Seed(time.Now().Unix())
	flag.Parse()
	logrus.SetFormatter(&logrus.JSONFormatter{})

	home, err := os.UserHomeDir()
	if err != nil {
		logrus.Error("cannot get home dir: ", err)
		return
	}

	cfg, err := config.New(home)
	if err != nil {
		logrus.Error("cannot load default configuration: ", err)
		return
	}

	db, err := store.New(cfg.DataDir, "keys", "contacts")
	if err != nil {
		logrus.Error("cannot start local database: ", err)
		return
	}
	defer db.Close()

	api := rpc.New(db, service.New())

	if cfg.RPCEnabled {
		logrus.Info("Starting RPC server at ", cfg.RPCListen)
		go func() {
			if err := rpc.ListenAndServe(api, cfg.RPCListen, cfg.RPCKey); err != nil {
				logrus.Error("error starting http json rpc server: ", err)
			}
		}()
	}
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	switch {
	case *noui:
		logrus.Info("No UI. Waiting for kill signal...")
		<-sigc
		logrus.Info("Exiting...")
	default:
		app := wails.CreateApp(&wails.AppConfig{
			Width:            900,
			Height:           1200,
			Title:            "ERA - Age encryption app",
			JS:               js,
			CSS:              css,
			HTML:             html,
			Resizable:        true,
			DisableInspector: !*debug,
		})
		binder := ui.New(cfg, api)
		app.Bind(binder)
		logrus.Info(wails.BuildMode)
		app.Run()
	}
}
