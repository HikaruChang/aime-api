package hub

import (
	"aime-api/config"
	"aime-api/hub/executor"
	"aime-api/hub/route"
)

type Option func(*config.General)

func Parse(options ...Option) error {
	cfg, err := executor.Parse()
	if err != nil {
		return err
	}

	for _, option := range options {
		option(cfg)
	}

	// Init server
	go route.Start(cfg)

	executor.ApplyConfig(cfg, true)
	return nil
}
