package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	config "aime-api/config"
	C "aime-api/constant"
	hub "aime-api/hub"
	log "aime-api/log"

	"go.uber.org/automaxprocs/maxprocs"
)

var (
	flagset    map[string]bool
	version    bool
	logFile    string
	configFile string
)

func init() {
	flag.StringVar(&configFile, "f", "", "specify configuration file")
	flag.StringVar(&logFile, "l", "", "specify log file")
	flag.BoolVar(&version, "v", false, "show current version of popnft")
	flag.Parse()

	flagset = map[string]bool{}
	flag.Visit(func(f *flag.Flag) {
		flagset[f.Name] = true
	})
}

func main() {
	maxprocs.Set(maxprocs.Logger(func(string, ...any) {}))

	if version {
		fmt.Printf(
			"AIME %s %s %s with %s %s\n",
			C.Version,
			runtime.GOOS,
			runtime.GOARCH,
			runtime.Version(),
			C.BuildTime,
		)
		return
	}

	if logFile != "" {
		if !filepath.IsAbs(logFile) {
			currentDir, _ := os.Getwd()
			logFile = filepath.Join(currentDir, logFile)
		}
		C.SetLog(logFile)
	} else {
		logFile = filepath.Join(C.Path.HomeDir(), C.Path.Log())
		C.SetLog(logFile)
	}

	if configFile != "" {
		if !filepath.IsAbs(configFile) {
			currentDir, _ := os.Getwd()
			configFile = filepath.Join(currentDir, configFile)
		}
		C.SetConfig(configFile)
	} else {
		configFile := filepath.Join(C.Path.HomeDir(), C.Path.Config())
		C.SetConfig(configFile)
	}

	if err := log.Init(C.Path.HomeDir()); err != nil {
		log.Fatalln("Initial configuration directory error: %s", err.Error())
	}

	if err := config.Init(C.Path.HomeDir()); err != nil {
		log.Fatalln("Initial log directory error: %s", err.Error())
	}

	var options []hub.Option
	if err := hub.Parse(options...); err != nil {
		log.Fatalln("Parse configuration error: %s", err.Error())
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
