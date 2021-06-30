package io

import (
	"bufio"
	"encoding/json"
	"fmt"
	"kubecfg/arguments"
	"kubecfg/config"
	"os"
)

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

func LoadConfigFromFileSystem() (config.Config, error) {
	var config config.Config
	configFile, err := os.Open(configFile())
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		return config, err
	}
	return config, err
}

func WriteUpdatedConfigToFileSystem(config *config.Config) int {
	configFile, err := os.OpenFile(configFile(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println("cannot open 'config.json'")
		return 1
	}
	defer configFile.Close()

	return writeToDisk(configFile, config)
}

func BackupKubeConfigFile() {
	if existingPath(kubectlConfigFile()) {
		if existingPath(kubectlConfigBackupFile()) {
			os.Remove(kubectlConfigBackupFile())
		}
		os.Rename(kubectlConfigFile(), kubectlConfigBackupFile())
	}
}

func RestoreBackedUpKubeConfigFile() {
	if existingPath(kubectlConfigBackupFile()) {
		os.Rename(kubectlConfigBackupFile(), kubectlConfigFile())
	}
}

func SaveCreatedContextFile(args *arguments.ContextAddArguments) {
	os.Rename(kubectlConfigFile(), projectContextFile(args.ProjectName, string(args.ContextName)))
}
