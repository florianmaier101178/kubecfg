package io

import (
	"bufio"
	"encoding/json"
	"fmt"
	"kubecfg/config"
	"os"
)

func SetupConfigDir() int {
	if existingDir(configDir()) {
		fmt.Println("config directory '~/.kubecfg' is existing and will be used")
		return 1
	}
	if _, err := os.Stat(configDir()); os.IsNotExist(err) {
		os.Mkdir(configDir(), 0755)
	}
	return 0
}

func WriteInitialConfigToFileSystem(config *config.Config) int {
	configFile, err := os.Create(configFile())
	if err != nil {
		fmt.Println("cannot create 'config.json'")
		return 1
	}
	defer configFile.Close()

	return writeToDisk(configFile, config)
}

func writeToDisk(file *os.File, config *config.Config) int {
	configJson, err := json.Marshal(config)
	writer := bufio.NewWriter(file)
	_, err = writer.Write(configJson)
	err = writer.Flush()

	if err != nil {
		fmt.Println("cannot write content to 'config.json'")
		return 1
	}

	return 0
}
