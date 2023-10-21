package executor

import (
	config "aime-api/config"
	C "aime-api/constant"
	"fmt"
	"os"
	"sync"
)

var mux sync.Mutex

func readConfig(path string) ([]byte, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("configuration file %s is empty", path)
	}

	return data, err
}

// Parse config with default config path
func Parse() (*config.General, error) {
	return ParseWithPath(C.Path.Config())
}

func ParseWithBytes(buf []byte) (*config.General, error) {
	return config.Parse(buf)
}

func ParseWithPath(path string) (*config.General, error) {
	mux.Lock()
	defer mux.Unlock()

	buf, err := readConfig(path)
	if err != nil {
		return nil, err
	}

	return ParseWithBytes(buf)
}

func ApplyConfig(cfg *config.General, force bool) {
	mux.Lock()
	defer mux.Unlock()
}
