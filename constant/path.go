package constant

import (
	"os"
	P "path"
)

const Name = "AIME"

type path struct {
	homeDir    string
	configFile string
	logFile    string
}

// Path is used to get the configuration path
var Path = func() *path {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir, _ = os.Getwd()
	}

	homeDir = P.Join(homeDir, ".config", Name)
	return &path{homeDir: homeDir, configFile: "config.json", logFile: "log.txt"}
}()

// SetConfig is used to set the configuration file
func SetConfig(file string) {
	Path.configFile = file
}

// SetLog is used to set the log file
func SetLog(file string) {
	Path.logFile = file
}

func (p *path) HomeDir() string {
	return p.homeDir
}

func (p *path) Log() string {
	return p.logFile
}

func (p *path) Config() string {
	return p.configFile
}
