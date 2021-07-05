package command

import (
	"fmt"
	cfg "kubecfg/config"
	"kubecfg/io"
)

func loadConfig() (*cfg.Config, int) {
	if io.IllegalConfigurationSetup() {
		fmt.Println("kubecfg is not properly configured")
		return &cfg.Config{}, 1
	}

	config, err := io.LoadConfigFromFileSystem()
	if err != nil {
		fmt.Println("could not load 'config.json'")
		return &cfg.Config{}, 1
	}

	return &config, 0
}
